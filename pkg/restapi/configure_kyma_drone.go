// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/joek/kyma-drone/pkg/drone"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/joek/kyma-drone/pkg/handlers"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

//go:generate swagger generate server --target ../../pkg --name KymaDrone --spec ../../api-docs.json

func configureFlags(api *operations.KymaDroneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KymaDroneAPI, drone drone.Drone) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HaltDroneHandler = handlers.NewPublicHaltDroneHandler(drone)
	api.LandDroneHandler = handlers.NewPublicLandDroneHandler(drone)
	api.LeftDroneHandler = handlers.NewPublicLeftDroneHandler(drone)
	api.RightDroneHandler = handlers.NewPublicRightDroneHandler(drone)
	api.ForwardDroneHandler = handlers.NewPublicForwardDroneHandler(drone)
	api.BackwardDroneHandler = handlers.NewPublicBackwardDroneHandler(drone)
	api.StartDroneHandler = handlers.NewPublicStartDroneHandler(drone)
	api.StopDroneHandler = handlers.NewPublicStopDroneHandler(drone)
	api.TakeOffDroneHandler = handlers.NewPublicTakeOffDroneHandler(drone)
	api.BackFlipDroneHandler = handlers.NewPublicBackFlipDroneHandler(drone)
	api.FrontFlipDroneHandler = handlers.NewPublicFrontFlipDroneHandler(drone)
	api.LeftFlipDroneHandler = handlers.NewPublicLeftFlipDroneHandler(drone)
	api.RightFlipDroneHandler = handlers.NewPublicRightFlipDroneHandler(drone)
	api.UpDroneHandler = handlers.NewPublicUpDroneHandler(drone)
	api.DownDroneHandler = handlers.NewPublicDownDroneHandler(drone)
	api.ClockwiseDroneHandler = handlers.NewPublicClockwiseDroneHandler(drone)
	api.CounterClockwiseDroneHandler = handlers.NewPublicCounterClockwiseDroneHandler(drone)
	api.TakePictureDroneHandler = handlers.NewPublicTakePictureDroneHandler(drone)
	api.EmergencyDroneHandler = handlers.NewPublicEmergencyDroneHandler(drone)
	api.FlatTrimDroneHandler = handlers.NewPublicFlatTrimDroneHandler(drone)
	api.LightControlDroneHandler = handlers.NewPublicLightControlDroneHandler(drone)
	api.GunControlDroneHandler = handlers.NewPublicGunControlDroneHandler(drone)
	api.ClawControlDroneHandler = handlers.NewPublicClawControlDroneHandler(drone)

	api.ServerShutdown = func() {
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
