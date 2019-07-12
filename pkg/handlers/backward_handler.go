package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicBackwardDroneHandler is handling request to Backward the dron
type PublicBackwardDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Backward drones
func (h PublicBackwardDroneHandler) Handle(params operations.BackwardDroneParams) middleware.Responder {
	err := h.drone.Backward(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewBackwardDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.BackwardDroneNoContent{}
}

// NewPublicBackwardDroneHandler is creating a new Backward Handler
func NewPublicBackwardDroneHandler(d drone.Drone) PublicBackwardDroneHandler {
	return PublicBackwardDroneHandler{d}
}
