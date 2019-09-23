package handlers

import (
	"fmt"
	"log"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicLandDroneHandler is handling request to Land the dron
type PublicLandDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Land drones
func (h PublicLandDroneHandler) Handle(params operations.LandDroneParams) middleware.Responder {
	log.Println("Landing")
	err := h.drone.Land()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewLandDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.LandDroneNoContent{}
}

// NewPublicLandDroneHandler is creating a new Land Handler
func NewPublicLandDroneHandler(d drone.Drone) PublicLandDroneHandler {
	return PublicLandDroneHandler{d}
}
