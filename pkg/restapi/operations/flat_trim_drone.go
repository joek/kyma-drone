// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FlatTrimDroneHandlerFunc turns a function with the right signature into a flat trim drone handler
type FlatTrimDroneHandlerFunc func(FlatTrimDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FlatTrimDroneHandlerFunc) Handle(params FlatTrimDroneParams) middleware.Responder {
	return fn(params)
}

// FlatTrimDroneHandler interface for that can handle valid flat trim drone params
type FlatTrimDroneHandler interface {
	Handle(FlatTrimDroneParams) middleware.Responder
}

// NewFlatTrimDrone creates a new http.Handler for the flat trim drone operation
func NewFlatTrimDrone(ctx *middleware.Context, handler FlatTrimDroneHandler) *FlatTrimDrone {
	return &FlatTrimDrone{Context: ctx, Handler: handler}
}

/*FlatTrimDrone swagger:route POST /flatTrim flatTrimDrone

FlatTrim tells the Minidrone to trim the sensors

*/
type FlatTrimDrone struct {
	Context *middleware.Context
	Handler FlatTrimDroneHandler
}

func (o *FlatTrimDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFlatTrimDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
