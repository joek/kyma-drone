package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicFrontFlipDroneHandler is handling request to  FrontFlip the dron
type PublicFrontFlipDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  FrontFlip drones
func (h PublicFrontFlipDroneHandler) Handle(params operations.FrontFlipDroneParams) middleware.Responder {
	err := h.drone.FrontFlip()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewFrontFlipDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.FrontFlipDroneNoContent{}
}

// NewPublicFrontFlipDroneHandler is creating a new  FrontFlip Handler
func NewPublicFrontFlipDroneHandler(d drone.Drone) PublicFrontFlipDroneHandler {
	return PublicFrontFlipDroneHandler{d}
}
