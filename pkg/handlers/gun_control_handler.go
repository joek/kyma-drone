package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicGunControlDroneHandler is handling request to GunControl the dron
type PublicGunControlDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to GunControl drones
func (h PublicGunControlDroneHandler) Handle(params operations.GunControlDroneParams) middleware.Responder {
	err := h.drone.GunControl(0)
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewGunControlDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.GunControlDroneNoContent{}
}

// NewPublicGunControlDroneHandler is creating a new GunControl Handler
func NewPublicGunControlDroneHandler(d drone.Drone) PublicGunControlDroneHandler {
	return PublicGunControlDroneHandler{d}
}
