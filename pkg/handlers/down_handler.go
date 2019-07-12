package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicDownDroneHandler is handling request to Down the dron
type PublicDownDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Down drones
func (h PublicDownDroneHandler) Handle(params operations.DownDroneParams) middleware.Responder {
	err := h.drone.Down(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewDownDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.DownDroneNoContent{}
}

// NewPublicDownDroneHandler is creating a new Down Handler
func NewPublicDownDroneHandler(d drone.Drone) PublicDownDroneHandler {
	return PublicDownDroneHandler{d}
}
