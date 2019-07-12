package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicEmergencyDroneHandler is handling request to  Emergency the dron
type PublicEmergencyDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  Emergency drones
func (h PublicEmergencyDroneHandler) Handle(params operations.EmergencyDroneParams) middleware.Responder {
	err := h.drone.Emergency()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewEmergencyDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.EmergencyDroneNoContent{}
}

// NewPublicEmergencyDroneHandler is creating a new  Emergency Handler
func NewPublicEmergencyDroneHandler(d drone.Drone) PublicEmergencyDroneHandler {
	return PublicEmergencyDroneHandler{d}
}
