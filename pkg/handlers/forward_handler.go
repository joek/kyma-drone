package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicForwardDroneHandler is handling request to Forward the dron
type PublicForwardDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Forward drones
func (h PublicForwardDroneHandler) Handle(params operations.ForwardDroneParams) middleware.Responder {
	err := h.drone.Forward(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewForwardDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.ForwardDroneNoContent{}
}

// NewPublicForwardDroneHandler is creating a new Forward Handler
func NewPublicForwardDroneHandler(d drone.Drone) PublicForwardDroneHandler {
	return PublicForwardDroneHandler{d}
}
