package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicRightDroneHandler is handling request to Right the dron
type PublicRightDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Right drones
func (h PublicRightDroneHandler) Handle(params operations.RightDroneParams) middleware.Responder {
	err := h.drone.Right(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewRightDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.RightDroneNoContent{}
}

// NewPublicRightDroneHandler is creating a new Right Handler
func NewPublicRightDroneHandler(d drone.Drone) PublicRightDroneHandler {
	return PublicRightDroneHandler{d}
}
