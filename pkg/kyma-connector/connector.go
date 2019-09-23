package connector

import (
	"bytes"
	"crypto/tls"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/tehcyx/kyma-integration/pkg/kyma/certificate"
)

// NewKymaConnector is creating a new connector instance and load the relevant information out of the config file
func NewKymaConnector(c string) (*KymaConnector, error) {
	err := os.MkdirAll(c, 0755)
	if err != nil {
		return nil, err
	}

	con := &KymaConnector{
		configFilePath: c,
		ca:             &certificate.CACertificate{},
	}

	con.ReadConfig()

	return con, nil
}

// WriteService is storing the service configuration to disc
func (c *KymaConnector) WriteService(s *Service) error {
	f := c.ServicePath(s.id)
	log.Println(f)
	b, err := json.Marshal(s)
	if err != nil {
		log.Println("Failed to generate json struct.")
		return err
	}

	err = ioutil.WriteFile(f, b, 0644)
	if err != nil {
		log.Println("couldn't write service config file")
		return err
	}

	return nil
}

// WriteConfig is storing the configuration on disc
func (c *KymaConnector) WriteConfig() error {
	f := c.ConfigPath()
	b, err := json.Marshal(c.AppInfo)
	if err != nil {
		log.Println("Failed to generate json struct.")
		return err
	}

	err = ioutil.WriteFile(f, b, 0644)
	if err != nil {
		log.Println("couldn't write config file")
		return err
	}
	return nil
}

// ReadConfig is reading the configuration from disc
func (c *KymaConnector) ReadConfig() error {
	f := c.ConfigPath()
	_, err := os.Stat(f)
	if err != nil {
		log.Println("No config available")
		return err
	}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println("Failed to read file")
		return err
	}

	appInfo := &certificate.ApplicationConnectResponse{}

	err = json.Unmarshal(b, appInfo)
	if err != nil {
		log.Println("Failed to parse json")
		return err
	}

	c.AppInfo = appInfo

	kf := c.PrivateKeyPath()
	_, err = os.Stat(kf)
	if err != nil {
		log.Println("No key available")
		return err
	}

	k, err := ioutil.ReadFile(kf)
	if err != nil {
		log.Println("Failed to read key file")
		return err
	}

	c.ca.PrivateKey = string(k[:])

	cf := c.PublicKeyPath()
	_, err = os.Stat(cf)
	if err != nil {
		log.Println("No cert available")
		return err
	}

	cert, err := ioutil.ReadFile(cf)
	if err != nil {
		log.Println("Failed to read cert file")
		return err
	}

	c.ca.PublicKey = string(cert[:])

	return nil
}

// Connect is retrieving the connector metadata and exchanging certificates with kyma.
func (c *KymaConnector) Connect(urlString string) (err error) {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Println("Error: need valid url")
		return err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		log.Println("Failed to call connection endpoint")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %d, Message: %s", resp.StatusCode, "error")
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	appData := &certificate.ApplicationConnectResponse{}

	unmarshalInfoErr := json.Unmarshal([]byte(bodyString), appData)
	if unmarshalInfoErr != nil {
		log.Println("could not parse response")
		return err
	}

	c.AppInfo = appData

	c.GenerateKeysAndCertificate(appData.Certificate.Subject)

	var jsonStr = []byte(fmt.Sprintf("{\"csr\":\"%s\"}", base64.StdEncoding.EncodeToString([]byte(c.ca.Csr))))
	csrResp, err := http.Post(appData.CsrURL, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("Could not send CSR")
		return err
	}

	defer csrResp.Body.Close()
	csrBodyBytes, _ := ioutil.ReadAll(csrResp.Body)
	csrBodyString := string(csrBodyBytes)

	certData := &certificate.CertConnectResponse{}

	unmarshalCertErr := json.Unmarshal([]byte(csrBodyString), certData)
	if unmarshalCertErr != nil {
		log.Println("could not parse CSR response")
		return err
	}

	decodedCert, decodeErr := base64.StdEncoding.DecodeString(certData.Cert)
	if decodeErr != nil {
		log.Printf("something went wrong decoding the response")
		return err
	}

	certData.Cert = string(decodedCert)
	certBytes := []byte(certData.Cert)

	errCert := ioutil.WriteFile(c.PublicKeyPath(), certBytes, 0644)
	if errCert != nil {
		log.Fatalf("couldn't write server cert: %s", errCert)
	}

	return c.WriteConfig()
}

// SendEvent is sending an event to kyma
func (c *KymaConnector) SendEvent(message json.RawMessage, eventType string, version string) (err error) {
	client, err := c.GetSecureClient()
	if err != nil {
		return
	}

	e := &Event{
		Type:        eventType,
		TypeVersion: version,
		ID:          uuid.New().String(),
		Time:        time.Now(),
		Data:        message,
	}

	b, err := json.Marshal(e)
	if err != nil {
		log.Println("Failed to generate event.")
		return err
	}

	resp, err := client.Post(c.AppInfo.API.EventsURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode != 200 {
		log.Printf("Failed to send event (Status: %d", resp.StatusCode)
		return errors.New(bodyString)
	}

	log.Println(bodyString)

	return
}

// Register is registering an service to kyma
func (c *KymaConnector) Register(apiDocs string, eventDocs string, serviceConfig string) (err error) {
	serviceDescription := new(Service)

	// Documentation part of the serviceDescription broken: https://github.com/kyma-project/kyma/issues/3347
	serviceDescription.Documentation = new(ServiceDocumentation)
	serviceDescription.Documentation.DisplayName = "Test"
	serviceDescription.Documentation.Description = "test decsription"
	serviceDescription.Documentation.Tags = []string{"Tag1", "Tag2"}
	serviceDescription.Documentation.Type = "Test Type"

	serviceDescription.Description = "API Description"
	serviceDescription.ShortDescription = "API Short Description"

	serviceDescription.Provider = "Kyma example"
	serviceDescription.Name = "Kyma example service"

	if serviceConfig != "" {
		log.Println("Read Service Config")
		err := c.ReadService(serviceConfig, serviceDescription)
		if err != nil {
			log.Printf("Failed to read service config: %s", serviceConfig)
			return err
		}
	}

	if apiDocs != "" {
		if serviceDescription.API == nil {
			log.Println("No Servic Description")
			serviceDescription.API = new(ServiceAPI)
			serviceDescription.API.TargetURL = "http://localhost:8080/"
		}

		serviceDescription.API.Spec, err = c.getApiDocs(apiDocs)
		if err != nil {
			return err
		}

	}

	if eventDocs != "" {

		serviceDescription.Events = new(ServiceEvent)

		serviceDescription.Events.Spec, err = c.getEvetDocs(eventDocs)
		if err != nil {
			return err
		}

	}

	jsonBytes, err := json.Marshal(serviceDescription)
	if err != nil {
		log.Printf("JSON marshal failed: %s", err)
		return
	}

	if c.AppInfo == nil || c.AppInfo.API.MetadataURL == "" {
		log.Printf("%s", fmt.Errorf("metadata url is missing, cannot proceed"))
		return
	}

	client, err := c.GetSecureClient()
	if err != nil {
		return err
	}

	resp, err := client.Post(c.AppInfo.API.MetadataURL, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("Couldn't register service: %s", err)
		return
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		log.Printf("could not dump response: %v", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("Successfully registered service with")
		log.Printf(bodyString)
	} else {
		log.Printf("Status: %d >%s< \n on URL: %s", resp.StatusCode, bodyString, c.AppInfo.API.MetadataURL)
		return errors.New("Failed to register")
	}

	id := &struct {
		ID string `json: "id"`
	}{}

	err = json.Unmarshal(bodyBytes, id)
	if err != nil {
		log.Println("Failed to parse registration response")
		return err
	}

	log.Printf("%v", id)
	serviceDescription.id = id.ID

	return c.WriteService(serviceDescription)
}

// Update is updating an exsisting service using the service ID
func (c *KymaConnector) Update(id string, apiDocs string, eventDocs string) (err error) {
	serviceDescription := new(Service)
	err = c.ReadService(c.ServicePath(id), serviceDescription)
	if err != nil {
		log.Printf("Failed to read service config: %s", c.ServicePath(id))
		return err
	}

	if apiDocs != "" {
		if serviceDescription.API == nil {
			serviceDescription.API = new(ServiceAPI)
			serviceDescription.API.TargetURL = "http://localhost:8080/"
		}

		serviceDescription.API.Spec, err = c.getApiDocs(apiDocs)
		if err != nil {
			return err
		}

	}

	if eventDocs != "" {

		serviceDescription.Events = new(ServiceEvent)

		serviceDescription.Events.Spec, err = c.getEvetDocs(eventDocs)
		if err != nil {
			return err
		}

	}

	jsonBytes, err := json.Marshal(serviceDescription)
	if err != nil {
		log.Printf("JSON marshal failed: %s", err)
		return
	}

	if c.AppInfo == nil || c.AppInfo.API.MetadataURL == "" {
		log.Printf("%s", fmt.Errorf("metadata url is missing, cannot proceed"))
		return
	}

	client, err := c.GetSecureClient()
	if err != nil {
		return err
	}

	url := c.AppInfo.API.MetadataURL + "/" + id
	log.Println(string(jsonBytes[:]))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Couldn't register service: %s", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("Successfully registered service with")
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Printf("Status: %d >%s<\n on URL: %s", resp.StatusCode, bodyString, url)
		return errors.New("Failed to Update")
	}

	return
}

// ReadService is loading a service description from disk
func (c *KymaConnector) ReadService(path string, s *Service) error {
	_, err := os.Stat(path)
	if err != nil {
		log.Println("No service config available")
		return err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Failed to read file")
		return err
	}

	err = json.Unmarshal(b, s)
	if err != nil {
		log.Println("Failed to parse json")
		return err
	}

	return nil
}

// GetSecureClient is returning an http client with client certificate enabled
func (c *KymaConnector) GetSecureClient() (*http.Client, error) {
	cert, err := tls.X509KeyPair([]byte(c.ca.PublicKey), []byte(c.ca.PrivateKey))
	if err != nil {
		log.Println("Can't load certificates")
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return &http.Client{Transport: transport}, nil

}

// GetCACertificate is loading application Certificate and keys from disk
func (c *KymaConnector) GetCACertificate() (ca *certificate.CACertificate, err error) {

	if c.ca.PrivateKey != "" && c.ca.PublicKey != "" && c.ca.Csr != "" {
		return c.ca, nil
	}

	_, errCSR := os.Stat(c.CsrPath())
	_, errPub := os.Stat(c.PublicKeyPath())
	_, errPriv := os.Stat(c.PrivateKeyPath())

	// read cert.csr
	if errCSR == nil && errPub == nil && errPriv == nil {
		csrBytes, err := ioutil.ReadFile(c.CsrPath())
		if err != nil {
			log.Fatal("Read error on csr file")
		}
		c.ca.Csr = string(csrBytes[:])
		pubKeyBytes, err := ioutil.ReadFile(c.PublicKeyPath())
		if err != nil {
			log.Fatal("Read error on pub file")
		}
		c.ca.PublicKey = string(pubKeyBytes[:])
		privKeyBytes, err := ioutil.ReadFile(c.PrivateKeyPath())
		if err != nil {
			log.Fatal("Read error on priv file")
		}
		c.ca.PrivateKey = string(privKeyBytes[:])
	}
	return ca, err
}

func (c *KymaConnector) ServicePath(id string) string {
	return path.Join(c.configFilePath, id+".json")
}

func (c *KymaConnector) ConfigPath() string {
	return path.Join(c.configFilePath, "config.json")
}

// CsrPath is proividing the csr file path
func (c *KymaConnector) CsrPath() string {
	return path.Join(c.configFilePath, "request.csr")
}

// PrivateKeyPath is proividing the private key file path
func (c *KymaConnector) PrivateKeyPath() string {
	return path.Join(c.configFilePath, "client.key")
}

// PublicKeyPath is proividing the public key file path
func (c *KymaConnector) PublicKeyPath() string {
	return path.Join(c.configFilePath, "client.crt")
}

// GenerateKeysAndCertificate generates keys and certificates
func (c *KymaConnector) GenerateKeysAndCertificate(subject string) {
	location := "Walldorf"
	province := "Walldorf"
	country := "DE"
	organization := "Organization"
	organizationalUnit := "OrgUnit"
	commonName := "api-test"

	if subject != "" {
		//TODO: add a more generic version of this, as it panics if the order of the elements in the subject line is changed
		subjectMatch := regexp.MustCompile("^O=(?P<o>.*),OU=(?P<ou>.*),L=(?P<l>.*),ST=(?P<st>.*),C=(?P<c>.*),CN=(?P<cn>.*)$")
		match := subjectMatch.FindStringSubmatch(subject)
		result := make(map[string]string)
		for i, name := range subjectMatch.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		location = result["l"]
		province = result["st"]
		country = result["c"]
		organization = result["o"]
		organizationalUnit = result["ou"]
		commonName = result["cn"]
		c.AppName = commonName
	}

	pkixName := pkix.Name{
		Locality:           []string{location},
		Province:           []string{province},
		Country:            []string{country},
		Organization:       []string{organization},
		OrganizationalUnit: []string{organizationalUnit},
		CommonName:         commonName,
	}
	genCert, err := certificate.GenerateCSR(pkixName, time.Duration(1200), 2048)
	if err != nil {
		log.Println(err)
	}
	//write files here
	csrBytes := []byte(genCert.Csr)
	pubKeyBytes := []byte(genCert.PublicKey)
	privKeyBytes := []byte(genCert.PrivateKey)
	errCSR := ioutil.WriteFile(c.CsrPath(), csrBytes, 0644)
	if errCSR != nil {
		log.Fatal("couldn't write csr")
	}
	errPub := ioutil.WriteFile(c.PublicKeyPath(), pubKeyBytes, 0644)
	if errPub != nil {
		log.Fatal("couldn't write pub key")
	}
	errPriv := ioutil.WriteFile(c.PrivateKeyPath(), privKeyBytes, 0644)
	if errPriv != nil {
		log.Fatal("couldn't write priv key")
	}

	c.ca = genCert
}

func (x *KymaConnector) getApiDocs(apiDocs string) (m json.RawMessage, err error) {
	log.Println("Load API Docs")
	apiBytes, err := ioutil.ReadFile(apiDocs)
	if err != nil {
		log.Println("Read error on API Docs")
		return
	}
	m = json.RawMessage(string(apiBytes[:]))
	return
}

func (x *KymaConnector) getEvetDocs(eventDocs string) (m json.RawMessage, err error) {
	log.Println("Load Event logs")
	eventsBytes, err := ioutil.ReadFile(eventDocs)
	if err != nil {
		log.Println("Read error on Event Docs")
		return
	}
	m = json.RawMessage(string(eventsBytes[:]))
	return
}
