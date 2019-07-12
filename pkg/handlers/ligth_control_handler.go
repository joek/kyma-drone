package handlers

import (
	"fmt"

	"gobot.io/x/gobot/platforms/parrot/minidrone"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicLightControlDroneHandler is handling request to LightControl the dron
type PublicLightControlDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to LightControl drones
func (h PublicLightControlDroneHandler) Handle(params operations.LightControlDroneParams) middleware.Responder {
	var m uint8
	switch *params.Value.Mode {
	case "LightFixed":
		m = minidrone.LightFixed
	case "LightBlinked":
		m = minidrone.LightBlinked
	case "LightOscillated":
		m = minidrone.LightOscillated
	default:
		c := int32(30)
		m := "Error: Unknown Light mode"
		er := operations.NewLightControlDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})

		return er
	}
	err := h.drone.LightControl(0, m, uint8(*params.Value.Intensity))
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewLightControlDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.LightControlDroneNoContent{}
}

// NewPublicLightControlDroneHandler is creating a new LightControl Handler
func NewPublicLightControlDroneHandler(d drone.Drone) PublicLightControlDroneHandler {
	return PublicLightControlDroneHandler{d}
}
