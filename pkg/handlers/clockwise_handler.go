package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicClockwiseDroneHandler is handling request to Clockwise the dron
type PublicClockwiseDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Clockwise drones
func (h PublicClockwiseDroneHandler) Handle(params operations.ClockwiseDroneParams) middleware.Responder {
	err := h.drone.Clockwise(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewClockwiseDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.ClockwiseDroneNoContent{}
}

// NewPublicClockwiseDroneHandler is creating a new Clockwise Handler
func NewPublicClockwiseDroneHandler(d drone.Drone) PublicClockwiseDroneHandler {
	return PublicClockwiseDroneHandler{d}
}
