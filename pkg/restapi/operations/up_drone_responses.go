// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/joek/kyma-drone/pkg/models"
)

// UpDroneNoContentCode is the HTTP code returned for type UpDroneNoContent
const UpDroneNoContentCode int = 204

/*UpDroneNoContent Drone is turning up

swagger:response upDroneNoContent
*/
type UpDroneNoContent struct {
}

// NewUpDroneNoContent creates UpDroneNoContent with default headers values
func NewUpDroneNoContent() *UpDroneNoContent {

	return &UpDroneNoContent{}
}

// WriteResponse to the client
func (o *UpDroneNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*UpDroneDefault unexpected error

swagger:response upDroneDefault
*/
type UpDroneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUpDroneDefault creates UpDroneDefault with default headers values
func NewUpDroneDefault(code int) *UpDroneDefault {
	if code <= 0 {
		code = 500
	}

	return &UpDroneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the up drone default response
func (o *UpDroneDefault) WithStatusCode(code int) *UpDroneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the up drone default response
func (o *UpDroneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the up drone default response
func (o *UpDroneDefault) WithPayload(payload *models.ErrorModel) *UpDroneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the up drone default response
func (o *UpDroneDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpDroneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
