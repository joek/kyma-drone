// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TakeOffDroneHandlerFunc turns a function with the right signature into a take off drone handler
type TakeOffDroneHandlerFunc func(TakeOffDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TakeOffDroneHandlerFunc) Handle(params TakeOffDroneParams) middleware.Responder {
	return fn(params)
}

// TakeOffDroneHandler interface for that can handle valid take off drone params
type TakeOffDroneHandler interface {
	Handle(TakeOffDroneParams) middleware.Responder
}

// NewTakeOffDrone creates a new http.Handler for the take off drone operation
func NewTakeOffDrone(ctx *middleware.Context, handler TakeOffDroneHandler) *TakeOffDrone {
	return &TakeOffDrone{Context: ctx, Handler: handler}
}

/*TakeOffDrone swagger:route POST /takeOff takeOffDrone

TakeOff tells the Minidrone to takeoff

*/
type TakeOffDrone struct {
	Context *middleware.Context
	Handler TakeOffDroneHandler
}

func (o *TakeOffDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTakeOffDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
