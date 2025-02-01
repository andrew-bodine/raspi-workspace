package temperature_monitor

import (
	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
)

//go:generate counterfeiter -o fakes/fake_ds18b20_sensor.go --fake-name FakeDS18B20Sensor . DS18B20Sensor
type DS18B20Sensor interface {
	Read() *DS18B20Result
}

type DS18B20 struct {
	Pin monitors.GoRaspberryPiIOPin
}

type DS18B20Result struct {
	Error       Error
	Message     string
	Temperature float64
}

type Error string

const (
	ErrorNoError     = "NoError"
)

func (ds18b20 *DS18B20) Read() *DS18B20Result {
	return nil
}

