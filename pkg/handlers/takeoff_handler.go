package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicTakeOffDroneHandler is handling request to  TakeOff the dron
type PublicTakeOffDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  TakeOff drones
func (h PublicTakeOffDroneHandler) Handle(params operations.TakeOffDroneParams) middleware.Responder {
	err := h.drone.TakeOff()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewTakeOffDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.TakeOffDroneNoContent{}
}

// NewPublicTakeOffDroneHandler is creating a new  TakeOff Handler
func NewPublicTakeOffDroneHandler(d drone.Drone) PublicTakeOffDroneHandler {
	return PublicTakeOffDroneHandler{d}
}
