package drone

import (
	"encoding/json"
	"log"

	connector "github.com/joek/kyma-drone/pkg/kyma-connector"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

// Drone interface is a generic drone
type Drone interface {
	Connection() gobot.Connection
	Name() string
	SetName(n string)
	Start() (err error)
	Halt() (err error)
	Init() (err error)
	GenerateAllStates() (err error)
	TakeOff() (err error)
	Land() (err error)
	FlatTrim() (err error)
	Emergency() (err error)
	TakePicture() (err error)
	StartPcmd()
	Up(val int) error
	Down(val int) error
	Forward(val int) error
	Backward(val int) error
	Right(val int) error
	Left(val int) error
	Clockwise(val int) error
	CounterClockwise(val int) error
	Stop() error
	StartRecording() error
	StopRecording() error
	HullProtection(protect bool) error
	Outdoor(outdoor bool) error
	FrontFlip() (err error)
	BackFlip() (err error)
	RightFlip() (err error)
	LeftFlip() (err error)
	LightControl(id uint8, mode uint8, intensity uint8) error
	ClawControl(id uint8, mode uint8) (err error)
	GunControl(id uint8) (err error)
	gobot.Eventer
}

func GetDroneWorker(d Drone, conn *connector.KymaConnector) func() {
	d.On(minidrone.Takeoff, func(data interface{}) {
		log.Println("Takeoff")
		conn.SendEvent(json.RawMessage(`{"drone": "`+d.Name()+`"}`), "drone.takeOff", "v1")
	})
	return func() {
	}
}
