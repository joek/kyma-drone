package handlers

import (
	"fmt"

	"gobot.io/x/gobot/platforms/parrot/minidrone"

	"github.com/joek/kyma-drone/pkg/drone"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicClawControlDroneHandler is handling request to ClawControl the dron
type PublicClawControlDroneHandler struct {
	drone drone.Drone
}

// Handle http Handler to ClawControl drones
func (h PublicClawControlDroneHandler) Handle(params operations.ClawControlDroneParams) middleware.Responder {
	var m uint8
	switch *params.Value.Mode {
	case "ClawOpen":
		m = minidrone.ClawOpen
	case "ClawClosed":
		m = minidrone.ClawClosed
	default:
		c := int32(30)
		m := "Error: Unknown Claw mode"
		er := operations.NewClawControlDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}
	err := h.drone.ClawControl(0, m)
	if err != nil {
		c := int32(20)
		m := fmt.Sprintf("Error: %s", err)
		er := operations.NewClawControlDroneDefault(-10)
		er.SetPayload(&models.ErrorModel{
			Code:    &c,
			Message: &m,
		})
		return er
	}

	return &operations.ClawControlDroneNoContent{}
}

// NewPublicClawControlDroneHandler is creating a new ClawControl Handler
func NewPublicClawControlDroneHandler(d drone.Drone) PublicClawControlDroneHandler {
	return PublicClawControlDroneHandler{d}
}
