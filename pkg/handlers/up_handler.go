package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicUpDroneHandler is handling request to Up the dron
type PublicUpDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to Up drones
func (h PublicUpDroneHandler) Handle(params operations.UpDroneParams) middleware.Responder {
	err := h.drone.Up(int(*params.Value.Value))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewUpDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.UpDroneNoContent{}
}

// NewPublicUpDroneHandler is creating a new Up Handler
func NewPublicUpDroneHandler(d drone.Drone) PublicUpDroneHandler {
	return PublicUpDroneHandler{d}
}
