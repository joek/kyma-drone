package drone

import (
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
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
	StartRobot() error
	StopRobot() error
}

// TestDrone is creating a Test Drone
type TestDrone struct {
	*TestDriver
	Robot
}

// MiniDrone is creating a Minit Drone
type MiniDrone struct {
	bleAdaptor *ble.ClientAdaptor
	*minidrone.Driver
	Robot
}

// Robot is enhancing the Gobot Robot
type Robot struct {
	*gobot.Robot
}

// NewDrone is creating a new Drone Configuration
func NewDrone(test bool) Drone {

	if test {
		d := &TestDrone{}
		d.TestDriver = NewTestDriver()

		d.Robot.Robot = gobot.NewRobot("minidrone",
			[]gobot.Connection{},
			[]gobot.Device{d},
		)
		return d
	}

	d := &MiniDrone{}
	d.bleAdaptor = ble.NewClientAdaptor(os.Getenv("DRONE_NAME"))
	d.Driver = minidrone.NewDriver(d.bleAdaptor)

	d.Robot.Robot = gobot.NewRobot("minidrone",
		[]gobot.Connection{d.bleAdaptor},
		[]gobot.Device{d},
	)
	return d
}

// StartRobot is starting the Robot
func (d *Robot) StartRobot() error {
	return d.Robot.Start()
}

// StopRobot is stopping the Robot
func (d *Robot) StopRobot() error {
	return d.Robot.Stop()
}
