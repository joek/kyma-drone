package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicStopDroneHandler is handling request to  Stop the dron
type PublicStopDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  Stop drones
func (h PublicStopDroneHandler) Handle(params operations.StopDroneParams) middleware.Responder {
	err := h.drone.Stop()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewStopDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.StopDroneNoContent{}
}

// NewPublicStopDroneHandler is creating a new  Stop Handler
func NewPublicStopDroneHandler(d drone.Drone) PublicStopDroneHandler {
	return PublicStopDroneHandler{d}
}
