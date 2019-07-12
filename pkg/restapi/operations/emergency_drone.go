// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// EmergencyDroneHandlerFunc turns a function with the right signature into a emergency drone handler
type EmergencyDroneHandlerFunc func(EmergencyDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn EmergencyDroneHandlerFunc) Handle(params EmergencyDroneParams) middleware.Responder {
	return fn(params)
}

// EmergencyDroneHandler interface for that can handle valid emergency drone params
type EmergencyDroneHandler interface {
	Handle(EmergencyDroneParams) middleware.Responder
}

// NewEmergencyDrone creates a new http.Handler for the emergency drone operation
func NewEmergencyDrone(ctx *middleware.Context, handler EmergencyDroneHandler) *EmergencyDrone {
	return &EmergencyDrone{Context: ctx, Handler: handler}
}

/*EmergencyDrone swagger:route POST /emergency emergencyDrone

Emergency tells the Minidrone to perform an emergency shutdown

*/
type EmergencyDrone struct {
	Context *middleware.Context
	Handler EmergencyDroneHandler
}

func (o *EmergencyDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewEmergencyDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
