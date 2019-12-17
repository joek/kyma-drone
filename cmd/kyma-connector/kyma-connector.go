package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joek/kyma-drone/pkg/kyma-connector"

	"github.com/jessevdk/go-flags"
)

var options Options

type Options struct {
	ConfigPath string `short:"c" long:"config" description:"Config File Path" `
}

type ConnectCommand struct {
	Url string `short:"u" long:"url" description:"Kyma Connector Url" required:"true"`
}

var connectCommand ConnectCommand

func (x *ConnectCommand) Execute(args []string) error {
	if options.ConfigPath == "" {
		options.ConfigPath = "./config"
	}

	c, err := connector.NewKymaConnector(options.ConfigPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = c.Connect(connectCommand.Url)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return nil
}

type RegisterCommand struct {
	APIDocs       string `short:"a" long:"api" description:"OpenAPI Spec File"`
	EventDocs     string `short:"e" long:"events" description:"Async API Spec File"`
	ServiceConfig string `short:"s" long:"service" description:"Service Spec File"`
}

var registerCommand RegisterCommand

func (x *RegisterCommand) Execute(args []string) error {
	c, err := connector.NewKymaConnector(options.ConfigPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = c.Register(x.APIDocs, x.EventDocs, x.ServiceConfig)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

type UpdateCommand struct {
	ID        string `short:"i" long:"id" description:"Service ID" required:"true"`
	APIDocs   string `short:"a" long:"api" description:"OpenAPI Spec File"`
	EventDocs string `short:"e" long:"events" description:"Async API Spec File"`
}

var updateCommand UpdateCommand

func (x *UpdateCommand) Execute(args []string) error {
	c, err := connector.NewKymaConnector(options.ConfigPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = c.Update(x.ID, x.APIDocs, x.EventDocs)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

type SendEvent struct {
	EventType    string `short:"t" long:"event-type" description:"Event Type"`
	EventVersion string `short:"v" long:"event-version" description:"Event Version e.g. v1"`
	Data         string `short:"d" long:"data" description:"Event Payload"`
}

var sendEvent SendEvent

func (x *SendEvent) Execute(args []string) error {
	c, err := connector.NewKymaConnector(options.ConfigPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = c.SendEvent(json.RawMessage([]byte(x.Data)), x.EventType, x.EventVersion)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

func main() {
	parser := flags.NewParser(&options, flags.Default)

	parser.AddCommand("register",
		"Register Services & Events",
		"Register Services & Events specifications to kyma.",
		&registerCommand)

	parser.AddCommand("connect",
		"Connect Application",
		"Connect Application to kyma.",
		&connectCommand)

	parser.AddCommand("update",
		"Update Service",
		"Update Service configuration.",
		&updateCommand)

	parser.AddCommand("event",
		"Send Test Event",
		"Send Test Event.",
		&sendEvent)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

}
