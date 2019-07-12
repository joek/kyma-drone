// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RightFlipDroneHandlerFunc turns a function with the right signature into a right flip drone handler
type RightFlipDroneHandlerFunc func(RightFlipDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RightFlipDroneHandlerFunc) Handle(params RightFlipDroneParams) middleware.Responder {
	return fn(params)
}

// RightFlipDroneHandler interface for that can handle valid right flip drone params
type RightFlipDroneHandler interface {
	Handle(RightFlipDroneParams) middleware.Responder
}

// NewRightFlipDrone creates a new http.Handler for the right flip drone operation
func NewRightFlipDrone(ctx *middleware.Context, handler RightFlipDroneHandler) *RightFlipDrone {
	return &RightFlipDrone{Context: ctx, Handler: handler}
}

/*RightFlipDrone swagger:route POST /rightFlip rightFlipDrone

RightFlip tells the Minidrone to do a Right Flip

*/
type RightFlipDrone struct {
	Context *middleware.Context
	Handler RightFlipDroneHandler
}

func (o *RightFlipDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRightFlipDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
