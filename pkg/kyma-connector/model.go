package connector

import (
	"encoding/json"
	"time"

	"github.com/tehcyx/kyma-integration/pkg/kyma/certificate"
)

// KymaConnector is managing the connection to the kyma system.
type KymaConnector struct {
	configFilePath string
	AppInfo        *certificate.ApplicationConnectResponse
	AppName        string
	ca             *certificate.CACertificate
}

// Service kyma service struct
type Service struct {
	id               string
	Provider         string                `json:"provider,omitempty"`
	Name             string                `json:"name,omitempty"`
	Description      string                `json:"description,omitempty"`
	ShortDescription string                `json:"shortDescription,omitempty"`
	Labels           *ServiceLabel         `json:"labels,omitempty"`
	API              *ServiceAPI           `json:"api,omitempty"`
	Events           *ServiceEvent         `json:"events,omitempty"`
	Documentation    *ServiceDocumentation `json:"documentation,omitempty"`
}

// ServiceLabel kyma service labels
type ServiceLabel map[string]string

// ServiceAPI kyma service api definition
type ServiceAPI struct {
	TargetURL   string              `json:"targetUrl,omitempty"`
	Spec        json.RawMessage     `json:"spec,omitempty"`
	Credentials *ServiceCredentials `json:"credentials,omitempty"`
}

// ServiceCredentials kyma service credentials definition
type ServiceCredentials struct {
	Basic *ServiceBasicCredentials `json:"basic,omitempty"`
	OAuth *ServiceOAuthCredentials `json:"oauth,omitempty"`
}

// ServiceBasicCredentials kyma basic auth service credentials
type ServiceBasicCredentials struct {
	ClientID string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// ServiceOAuthCredentials kyma oauth service credentials
type ServiceOAuthCredentials struct {
	ClientID     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	URL          string `json:"url,omitempty"`
}

// ServiceEvent kyma service event definition
type ServiceEvent struct {
	Spec json.RawMessage `json:"spec,omitempty"`
}

// ServiceDocumentation kyma service documentation definition
type ServiceDocumentation struct {
	DisplayName string                     `json:"displayName,omitempty"`
	Description string                     `json:"description,omitempty"`
	Type        string                     `json:"type,omitempty"`
	Tags        []string                   `json:"tags,omitempty"`
	Docs        []*ServiceDocumentationDoc `json:"docs,omitempty"`
}

// ServiceDocumentationDoc kyma service documentation doc definition
type ServiceDocumentationDoc struct {
	Title  string `json:"title,omitempty"`
	Type   string `json:"type,omitempty"`
	Source string `json:"source,omitempty"`
}

// Event kyma event
type Event struct {
	Type        string          `json:"event-type"`
	TypeVersion string          `json:"event-type-version"`
	ID          string          `json:"event-id,omitempty"`
	Time        time.Time       `json:"event-time,omitempty"`
	Data        json.RawMessage `json:"data"`
}
