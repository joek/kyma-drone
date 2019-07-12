package drone

import (
	"log"
	"time"

	"gobot.io/x/gobot"
)

type TestDriver struct {
	name       string
	connection gobot.Connection
	flying     bool
	gobot.Eventer
}

const (
	// Battery event
	Battery = "battery"

	// FlightStatus event
	FlightStatus = "flightstatus"

	// Takeoff event
	Takeoff = "takeoff"

	// Hovering event
	Hovering = "hovering"

	// Flying event
	Flying = "flying"

	// Landing event
	Landing = "landing"

	// Landed event
	Landed = "landed"

	// Emergency event
	Emergency = "emergency"

	// Rolling event
	Rolling = "rolling"

	// FlatTrimChange event
	FlatTrimChange = "flattrimchange"

	// LightFixed mode for LightControl
	LightFixed = 0

	// LightBlinked mode for LightControl
	LightBlinked = 1

	// LightOscillated mode for LightControl
	LightOscillated = 3

	// ClawOpen mode for ClawControl
	ClawOpen = 0

	// ClawClosed mode for ClawControl
	ClawClosed = 1
)

// NewTestDriver creates a Test Minidrone Driver
func NewTestDriver() *TestDriver {
	n := &TestDriver{
		name:    gobot.DefaultName("TestMinidrone"),
		Eventer: gobot.NewEventer(),
	}

	n.AddEvent(Battery)
	n.AddEvent(FlightStatus)

	n.AddEvent(Takeoff)
	n.AddEvent(Flying)
	n.AddEvent(Hovering)
	n.AddEvent(Landing)
	n.AddEvent(Landed)
	n.AddEvent(Emergency)
	n.AddEvent(Rolling)

	return n
}

// Connection returns the BLE connection
func (b *TestDriver) Connection() gobot.Connection { return b.connection }

// Name returns the Driver Name
func (b *TestDriver) Name() string { return b.name }

// SetName sets the Driver Name
func (b *TestDriver) SetName(n string) { b.name = n }

// Start tells driver to get ready to do work
func (b *TestDriver) Start() (err error) {
	log.Println("Start Drone")
	b.Init()

	return
}

// Halt stops minidrone driver (void)
func (b *TestDriver) Halt() (err error) {
	log.Println("Halt Drone")
	b.Land()

	time.Sleep(500 * time.Millisecond)
	return
}

// Init initializes the BLE insterfaces used by the Minidrone
func (b *TestDriver) Init() (err error) {
	log.Println("Init Drone")
	return
}

// GenerateAllStates sets up all the default states aka settings on the drone
func (b *TestDriver) GenerateAllStates() (err error) {
	return
}

// TakeOff tells the Minidrone to takeoff
func (b *TestDriver) TakeOff() (err error) {
	b.Publish(Takeoff, true)
	return
}

// Land tells the Minidrone to land
func (b *TestDriver) Land() (err error) {
	b.Publish(Landing, true)
	b.flying = false
	b.Publish(Landed, true)

	log.Println("Land drone")
	return err
}

// FlatTrim calibrates the Minidrone to use its current position as being level
func (b *TestDriver) FlatTrim() (err error) {
	return err
}

// Emergency sets the Minidrone into emergency mode
func (b *TestDriver) Emergency() (err error) {
	b.Publish(Emergency, true)
	return err
}

// TakePicture tells the Minidrone to take a picture
func (b *TestDriver) TakePicture() (err error) {
	return err
}

// StartPcmd starts the continuous Pcmd communication with the Minidrone
func (b *TestDriver) StartPcmd() {
}

// Up tells the drone to ascend. Pass in an int from 0-100.
func (b *TestDriver) Up(val int) error {
	return nil
}

// Down tells the drone to descend. Pass in an int from 0-100.
func (b *TestDriver) Down(val int) error {
	return nil
}

// Forward tells the drone to go forward. Pass in an int from 0-100.
func (b *TestDriver) Forward(val int) error {
	log.Printf("Forward: %d", val)

	return nil
}

// Backward tells drone to go in reverse. Pass in an int from 0-100.
func (b *TestDriver) Backward(val int) error {
	log.Printf("Backward: %d", val)

	return nil
}

// Right tells drone to go right. Pass in an int from 0-100.
func (b *TestDriver) Right(val int) error {
	log.Printf("Right: %d", val)

	return nil
}

// Left tells drone to go left. Pass in an int from 0-100.
func (b *TestDriver) Left(val int) error {
	log.Printf("Left: %d", val)
	return nil
}

// Clockwise tells drone to rotate in a clockwise direction. Pass in an int from 0-100.
func (b *TestDriver) Clockwise(val int) error {
	return nil
}

// CounterClockwise tells drone to rotate in a counter-clockwise direction.
// Pass in an int from 0-100.
func (b *TestDriver) CounterClockwise(val int) error {
	return nil
}

// Stop tells the drone to stop moving in any direction and simply hover in place
func (b *TestDriver) Stop() error {
	b.flying = true
	b.Publish(Hovering, true)
	return nil
}

// StartRecording is not supported by the Parrot Minidrone
func (b *TestDriver) StartRecording() error {
	return nil
}

// StopRecording is not supported by the Parrot Minidrone
func (b *TestDriver) StopRecording() error {
	return nil
}

// HullProtection is not supported by the Parrot Minidrone
func (b *TestDriver) HullProtection(protect bool) error {
	return nil
}

// Outdoor mode is not supported by the Parrot Minidrone
func (b *TestDriver) Outdoor(outdoor bool) error {
	return nil
}

// FrontFlip tells the drone to perform a front flip
func (b *TestDriver) FrontFlip() (err error) {
	b.Publish(Rolling, true)
	return nil
}

// BackFlip tells the drone to perform a backflip
func (b *TestDriver) BackFlip() (err error) {
	b.Publish(Rolling, true)
	return nil
}

// RightFlip tells the drone to perform a flip to the right
func (b *TestDriver) RightFlip() (err error) {
	b.Publish(Rolling, true)
	return nil
}

// LeftFlip tells the drone to perform a flip to the left
func (b *TestDriver) LeftFlip() (err error) {
	b.Publish(Rolling, true)
	return nil
}

// LightControl controls lights on those Minidrone models which
// have the correct hardware, such as the Maclane, Blaze, & Swat.
// Params:
//		id - always 0
//		mode - either LightFixed, LightBlinked, or LightOscillated
//		intensity - Light intensity from 0 (OFF) to 100 (Max intensity).
// 					Only used in LightFixed mode.
//
func (b *TestDriver) LightControl(id uint8, mode uint8, intensity uint8) (err error) {
	return
}

// ClawControl controls the claw on the Parrot Mambo
// Params:
//		id - always 0
//		mode - either ClawOpen or ClawClosed
//
func (b *TestDriver) ClawControl(id uint8, mode uint8) (err error) {
	log.Printf("Claw mode: %d", mode)
	return
}

// GunControl fires the gun on the Parrot Mambo
// Params:
//		id - always 0
//
func (b *TestDriver) GunControl(id uint8) (err error) {
	return
}
