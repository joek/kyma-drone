// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BackFlipDroneHandlerFunc turns a function with the right signature into a back flip drone handler
type BackFlipDroneHandlerFunc func(BackFlipDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BackFlipDroneHandlerFunc) Handle(params BackFlipDroneParams) middleware.Responder {
	return fn(params)
}

// BackFlipDroneHandler interface for that can handle valid back flip drone params
type BackFlipDroneHandler interface {
	Handle(BackFlipDroneParams) middleware.Responder
}

// NewBackFlipDrone creates a new http.Handler for the back flip drone operation
func NewBackFlipDrone(ctx *middleware.Context, handler BackFlipDroneHandler) *BackFlipDrone {
	return &BackFlipDrone{Context: ctx, Handler: handler}
}

/*BackFlipDrone swagger:route POST /backFlip backFlipDrone

BackFlip tells the Minidrone to do a Back Flip

*/
type BackFlipDrone struct {
	Context *middleware.Context
	Handler BackFlipDroneHandler
}

func (o *BackFlipDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBackFlipDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
