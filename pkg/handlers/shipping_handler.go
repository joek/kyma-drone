package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"gobot.io/x/gobot/platforms/parrot/minidrone"

	"github.com/joek/kyma-drone/pkg/drone"
	connector "github.com/joek/kyma-drone/pkg/kyma-connector"
	"github.com/joek/kyma-drone/pkg/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/joek/kyma-drone/pkg/restapi/operations"
)

// PublicShippingDroneHandler is handling request to ship the package
type PublicShippingDroneHandler struct {
	drone drone.Drone
	conn  *connector.KymaConnector
}

// Handle http Handler to Up drones
func (h PublicShippingDroneHandler) Handle(params operations.ShipPackageParams) middleware.Responder {
	log.Println("Ship Package")
	orderCoder := params.Value.OrderCode
	var mux sync.Mutex
	mux.Lock()
	defer mux.Unlock()

	h.drone.Once(minidrone.Landed, func(data interface{}) {
		log.Println("Landed")
		h.conn.SendEvent(json.RawMessage([]byte("{\"orderCode\": \""+*orderCoder+"\"}")), "drone.shipped", "v1")
	})
	h.drone.Once(minidrone.Hovering, func(data interface{}) {
		mux.Unlock()
	})
	err := h.drone.TakeOff()
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
	log.Println("Wait for Hovering")
	mux.Lock()
	log.Println("Hovering")

	return &operations.UpDroneNoContent{}
}

// NewPublicShippingDroneHandler is creating a new Up Handler
func NewPublicShippingDroneHandler(d drone.Drone, conn *connector.KymaConnector) PublicShippingDroneHandler {
	return PublicShippingDroneHandler{d, conn}
}
