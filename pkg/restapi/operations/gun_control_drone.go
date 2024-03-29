// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GunControlDroneHandlerFunc turns a function with the right signature into a gun control drone handler
type GunControlDroneHandlerFunc func(GunControlDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GunControlDroneHandlerFunc) Handle(params GunControlDroneParams) middleware.Responder {
	return fn(params)
}

// GunControlDroneHandler interface for that can handle valid gun control drone params
type GunControlDroneHandler interface {
	Handle(GunControlDroneParams) middleware.Responder
}

// NewGunControlDrone creates a new http.Handler for the gun control drone operation
func NewGunControlDrone(ctx *middleware.Context, handler GunControlDroneHandler) *GunControlDrone {
	return &GunControlDrone{Context: ctx, Handler: handler}
}

/*GunControlDrone swagger:route POST /gunControl gunControlDrone

GunControl tells the Minidrone to shoot

*/
type GunControlDrone struct {
	Context *middleware.Context
	Handler GunControlDroneHandler
}

func (o *GunControlDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGunControlDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
