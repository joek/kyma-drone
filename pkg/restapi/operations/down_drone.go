// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DownDroneHandlerFunc turns a function with the right signature into a down drone handler
type DownDroneHandlerFunc func(DownDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DownDroneHandlerFunc) Handle(params DownDroneParams) middleware.Responder {
	return fn(params)
}

// DownDroneHandler interface for that can handle valid down drone params
type DownDroneHandler interface {
	Handle(DownDroneParams) middleware.Responder
}

// NewDownDrone creates a new http.Handler for the down drone operation
func NewDownDrone(ctx *middleware.Context, handler DownDroneHandler) *DownDrone {
	return &DownDrone{Context: ctx, Handler: handler}
}

/*DownDrone swagger:route POST /down downDrone

Down tells drone to go down.

*/
type DownDrone struct {
	Context *middleware.Context
	Handler DownDroneHandler
}

func (o *DownDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDownDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
