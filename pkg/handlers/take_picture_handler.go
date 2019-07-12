package handlers

import (
	"fmt"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicTakePictureDroneHandler is handling request to  TakePicture the dron
type PublicTakePictureDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to  TakePicture drones
func (h PublicTakePictureDroneHandler) Handle(params operations.TakePictureDroneParams) middleware.Responder {
	err := h.drone.TakePicture()
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewTakePictureDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	return &operations.TakePictureDroneNoContent{}
}

// NewPublicTakePictureDroneHandler is creating a new  TakePicture Handler
func NewPublicTakePictureDroneHandler(d drone.Drone) PublicTakePictureDroneHandler {
	return PublicTakePictureDroneHandler{d}
}
