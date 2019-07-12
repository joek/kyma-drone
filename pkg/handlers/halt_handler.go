package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicHaltDroneHandler is handling request to  Halt the dron
type PublicHaltDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  Halt drones
func (h PublicHaltDroneHandler) Handle(params operations.HaltDroneParams) middleware.Responder {
	err := h.drone.Halt()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewHaltDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.HaltDroneNoContent{}
}

// NewPublicHaltDroneHandler is creating a new  Halt Handler
func NewPublicHaltDroneHandler(d drone.Drone) PublicHaltDroneHandler {
	return PublicHaltDroneHandler{d}
}
