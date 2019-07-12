// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/handlers"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

//go:generate swagger generate server --target ../../pkg --name KymaDrone --spec ../../api-docs.json

func configureFlags(api *operations.KymaDroneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KymaDroneAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	t := false
	if os.Getenv("TEST") == "true" {
		t = true
	}
	d := drone.NewDrone(t)

	go func() {
		err := d.StartRobot()
		if err != nil {
			log.Fatal("Failed to start Robot")
		}
	}()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HaltDroneHandler = handlers.NewPublicHaltDroneHandler(d)
	api.LandDroneHandler = handlers.NewPublicLandDroneHandler(d)
	api.LeftDroneHandler = handlers.NewPublicLeftDroneHandler(d)
	api.RightDroneHandler = handlers.NewPublicRightDroneHandler(d)
	api.ForwardDroneHandler = handlers.NewPublicForwardDroneHandler(d)
	api.BackwardDroneHandler = handlers.NewPublicBackwardDroneHandler(d)
	api.StartDroneHandler = handlers.NewPublicStartDroneHandler(d)
	api.StopDroneHandler = handlers.NewPublicStopDroneHandler(d)
	api.TakeOffDroneHandler = handlers.NewPublicTakeOffDroneHandler(d)
	api.BackFlipDroneHandler = handlers.NewPublicBackFlipDroneHandler(d)
	api.FrontFlipDroneHandler = handlers.NewPublicFrontFlipDroneHandler(d)
	api.LeftFlipDroneHandler = handlers.NewPublicLeftFlipDroneHandler(d)
	api.RightFlipDroneHandler = handlers.NewPublicRightFlipDroneHandler(d)
	api.UpDroneHandler = handlers.NewPublicUpDroneHandler(d)
	api.DownDroneHandler = handlers.NewPublicDownDroneHandler(d)
	api.ClockwiseDroneHandler = handlers.NewPublicClockwiseDroneHandler(d)
	api.CounterClockwiseDroneHandler = handlers.NewPublicCounterClockwiseDroneHandler(d)
	api.TakePictureDroneHandler = handlers.NewPublicTakePictureDroneHandler(d)
	api.EmergencyDroneHandler = handlers.NewPublicEmergencyDroneHandler(d)
	api.FlatTrimDroneHandler = handlers.NewPublicFlatTrimDroneHandler(d)
	api.LightControlDroneHandler = handlers.NewPublicLightControlDroneHandler(d)
	api.GunControlDroneHandler = handlers.NewPublicGunControlDroneHandler(d)
	api.ClawControlDroneHandler = handlers.NewPublicClawControlDroneHandler(d)

	api.ServerShutdown = func() {
		d.StopRobot()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
