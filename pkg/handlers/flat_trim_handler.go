package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicFlatTrimDroneHandler is handling request to  FlatTrim the dron
type PublicFlatTrimDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  FlatTrim drones
func (h PublicFlatTrimDroneHandler) Handle(params operations.FlatTrimDroneParams) middleware.Responder {
	err := h.drone.FlatTrim()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewFlatTrimDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.FlatTrimDroneNoContent{}
}

// NewPublicFlatTrimDroneHandler is creating a new  FlatTrim Handler
func NewPublicFlatTrimDroneHandler(d drone.Drone) PublicFlatTrimDroneHandler {
	return PublicFlatTrimDroneHandler{d}
}
