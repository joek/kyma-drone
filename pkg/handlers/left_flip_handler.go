package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicLeftFlipDroneHandler is handling request to  LeftFlip the dron
type PublicLeftFlipDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  LeftFlip drones
func (h PublicLeftFlipDroneHandler) Handle(params operations.LeftFlipDroneParams) middleware.Responder {
	err := h.drone.LeftFlip()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewLeftFlipDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.LeftFlipDroneNoContent{}
}

// NewPublicLeftFlipDroneHandler is creating a new  LeftFlip Handler
func NewPublicLeftFlipDroneHandler(d drone.Drone) PublicLeftFlipDroneHandler {
	return PublicLeftFlipDroneHandler{d}
}
