// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TakePictureDroneHandlerFunc turns a function with the right signature into a take picture drone handler
type TakePictureDroneHandlerFunc func(TakePictureDroneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TakePictureDroneHandlerFunc) Handle(params TakePictureDroneParams) middleware.Responder {
	return fn(params)
}

// TakePictureDroneHandler interface for that can handle valid take picture drone params
type TakePictureDroneHandler interface {
	Handle(TakePictureDroneParams) middleware.Responder
}

// NewTakePictureDrone creates a new http.Handler for the take picture drone operation
func NewTakePictureDrone(ctx *middleware.Context, handler TakePictureDroneHandler) *TakePictureDrone {
	return &TakePictureDrone{Context: ctx, Handler: handler}
}

/*TakePictureDrone swagger:route POST /takePicture takePictureDrone

TakePicture tells the Minidrone to take a picture

*/
type TakePictureDrone struct {
	Context *middleware.Context
	Handler TakePictureDroneHandler
}

func (o *TakePictureDrone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTakePictureDroneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
