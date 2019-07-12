package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicStartDroneHandler is handling request to start the dron
type PublicStartDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to start drones
func (h PublicStartDroneHandler) Handle(params operations.StartDroneParams) middleware.Responder {
	err := h.drone.Start()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewStartDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.StartDroneNoContent{}
}

// NewPublicStartDroneHandler is creating a new Start Handler
func NewPublicStartDroneHandler(d drone.Drone) PublicStartDroneHandler {
	return PublicStartDroneHandler{d}
}
