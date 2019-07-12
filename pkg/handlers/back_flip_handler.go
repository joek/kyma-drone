package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicBackFlipDroneHandler is handling request to  BackFlip the dron
type PublicBackFlipDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  BackFlip drones
func (h PublicBackFlipDroneHandler) Handle(params operations.BackFlipDroneParams) middleware.Responder {
	err := h.drone.BackFlip()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewBackFlipDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.BackFlipDroneNoContent{}
}

// NewPublicBackFlipDroneHandler is creating a new  BackFlip Handler
func NewPublicBackFlipDroneHandler(d drone.Drone) PublicBackFlipDroneHandler {
	return PublicBackFlipDroneHandler{d}
}
