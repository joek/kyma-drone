package main

import (
	"log"
	"os"

	"github.com/joek/kyma-drone/pkg/drone"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	connector "github.com/joek/kyma-drone/pkg/kyma-connector"
	"github.com/joek/kyma-drone/pkg/restapi"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

func main() {
	c, err := connector.NewKymaConnector(os.Getenv("KYMA_CONFIG"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var robot *gobot.Robot
	var driver drone.Drone

	if os.Getenv("TEST_API") == "true" {
		d := drone.NewTestDriver()
		driver = d

		robot = gobot.NewRobot("minidrone",
			[]gobot.Connection{},
			[]gobot.Device{driver},
			drone.GetDroneWorker(d, c),
		)
	} else {
		bleAdaptor := ble.NewClientAdaptor(os.Args[1])
		d := minidrone.NewDriver(bleAdaptor)
		driver = d

		robot = gobot.NewRobot("minidrone",
			[]gobot.Connection{bleAdaptor},
			[]gobot.Device{driver},
			drone.GetDroneWorker(d, c),
		)
	}

	go func() {
		robot.Start()
	}()
	defer robot.Stop()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewKymaDroneAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Kyma Drone"
	parser.LongDescription = "A simple API to controll remote controlled drones"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI(driver, c)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
