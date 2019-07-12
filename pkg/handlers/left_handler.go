package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicLeftDroneHandler is handling request to Left the dron
type PublicLeftDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Left drones
func (h PublicLeftDroneHandler) Handle(params operations.LeftDroneParams) middleware.Responder {
	err := h.drone.Left(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewLeftDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.LeftDroneNoContent{}
}

// NewPublicLeftDroneHandler is creating a new Left Handler
func NewPublicLeftDroneHandler(d drone.Drone) PublicLeftDroneHandler {
	return PublicLeftDroneHandler{d}
}
