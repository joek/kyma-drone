package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicCounterClockwiseDroneHandler is handling request to CounterClockwise the dron
type PublicCounterClockwiseDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to CounterClockwise drones
func (h PublicCounterClockwiseDroneHandler) Handle(params operations.CounterClockwiseDroneParams) middleware.Responder {
	err := h.drone.CounterClockwise(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewCounterClockwiseDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.CounterClockwiseDroneNoContent{}
}

// NewPublicCounterClockwiseDroneHandler is creating a new CounterClockwise Handler
func NewPublicCounterClockwiseDroneHandler(d drone.Drone) PublicCounterClockwiseDroneHandler {
	return PublicCounterClockwiseDroneHandler{d}
}
