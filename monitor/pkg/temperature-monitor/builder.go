package temperature_monitor

import (
	"flag"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
)

var (
	TempPinFlag         = flag.Int("tempPin", -1, "Raspberry Pi GPIO pin to monitor for temperature and humidity changes.")
	TempPinPollRateFlag = flag.Duration("tempPinPollRate", time.Second*5, "How often to check the state of --tempPin.")
)

type TemperatureMonitorConfig struct {
	Logger *zap.Logger

	// Represents the GPIO pin that this monitor is pay attention to.
	GPIOPin monitors.GoRaspberryPiIOPin

	GPIOPinPollRate time.Duration
}

// Implement the monitors.ConfigWithFlagsAndLoggerConstructor interface.
func NewConfigFromFlagsWithLogger(logger *zap.Logger) (interface{}, error) {
	if *TempPinFlag == -1 {
		return nil, nil
	}

	pin := rpio.Pin(*TempPinFlag)

	config := TemperatureMonitorConfig{
		Logger:          logger,
		GPIOPin:         pin,
		GPIOPinPollRate: *TempPinPollRateFlag,
	}

	return &config, nil
}
