package monitors

import (
	gorpio "github.com/stianeikeland/go-rpio"
)

// Aggregated Pin methods exposed by the gorpio package.
//go:generate counterfeiter -o fakes/fake_go_raspberry_pi_io_pin.go --fake-name FakeGoRaspberryPiIOPin . GoRaspberryPiIOPin
type GoRaspberryPiIOPin interface {
	Input()
	Output()
	Clock()
	Pwm()
	High()
	Low()
	Toggle()
	Freq(freq int)
	DutyCycle(dutyLen, cycleLen uint32)
	Mode(mode gorpio.Mode)
	Write(state gorpio.State)
	Read() gorpio.State
	Pull(pull gorpio.Pull)
	PullUp()
	PullOff()
	Detect(edge gorpio.Edge)
	EdgeDetected() bool
}

// Aggregated public functions exposed by the gorpio package.
//go:generate counterfeiter -o fakes/fake_go_raspberry_pi_io.go --fake-name FakeGoRaspberryPiIO . GoRaspberryPiIO
type GoRaspberryPiIO interface {
	PinMode(pin GoRaspberryPiIOPin, mode gorpio.Mode)
	WritePin(pin GoRaspberryPiIOPin, state gorpio.State)
	ReadPin(pin GoRaspberryPiIOPin) gorpio.State
	TogglePin(pin GoRaspberryPiIOPin)
	DetectEdge(pin GoRaspberryPiIOPin, edge gorpio.Edge)
	EdgeDetected(pin GoRaspberryPiIOPin) bool
	PullMode(pin GoRaspberryPiIOPin, pull gorpio.Pull)
	SetFreq(pin GoRaspberryPiIOPin, freq int)
	SetDutyCycle(pin GoRaspberryPiIOPin, dutyLen, cycleLen uint32)
	StopPwm()
	StartPwm()
	EnableIRQs(irqs uint64)
	DisableIRQs(irqs uint64)
	Open() (err error)
	Close() error
}
