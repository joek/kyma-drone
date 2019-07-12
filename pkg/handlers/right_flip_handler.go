package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicRightFlipDroneHandler is handling request to  RightFlip the dron
type PublicRightFlipDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  RightFlip drones
func (h PublicRightFlipDroneHandler) Handle(params operations.RightFlipDroneParams) middleware.Responder {
	err := h.drone.RightFlip()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewRightFlipDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.RightFlipDroneNoContent{}
}

// NewPublicRightFlipDroneHandler is creating a new  RightFlip Handler
func NewPublicRightFlipDroneHandler(d drone.Drone) PublicRightFlipDroneHandler {
	return PublicRightFlipDroneHandler{d}
}
