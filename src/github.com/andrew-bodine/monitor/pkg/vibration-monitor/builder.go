package vibration_monitor

import (
	"flag"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"

	"github.com/andrew-bodine/monitor/pkg/monitors"
)

var (
	VibrPinFlag         = flag.Int("vibrPin", -1, "Raspberry Pi GPIO pin to monitor for vibrations.")
	VibrPinPollRateFlag = flag.Duration("vibrPinPollRate", time.Millisecond*500, "How often to check the state of --vibrPin.")
	VibrMinPeriodFlag   = flag.Duration("vibrMinPeriod", time.Second*15, "Minimum period of sustained vibration that constitutes an event.")
)

// Collection of configurable settings on that are common amongst vibration
// monitor implementations.
type VibrationMonitorConfig struct {
	Logger *zap.Logger

	// Represents the GPIO pin that this monitor is pay attention to.
	GPIOPin monitors.GoRaspberryPiIOPin

	GPIOPinPollRate time.Duration

	MinVibrationPeriod time.Duration
}

// Implement the monitors.ConfigWithFlagsAndLoggerConstructor interface.
func NewConfigFromFlagsWithLogger(logger *zap.Logger) (interface{}, error) {
	if *VibrPinFlag == -1 {
		return nil, nil
		// return nil, errors.New("Required flag --vibrPin missing, please specify and try again.")
	}

	// Get a reference to the desired pin.
	pin := rpio.Pin(*VibrPinFlag)

	// Set pin in input mode.
	pin.Input()

	config := VibrationMonitorConfig{
		Logger:             logger,
		GPIOPin:            pin,
		GPIOPinPollRate:    *VibrPinPollRateFlag,
		MinVibrationPeriod: *VibrMinPeriodFlag,
	}

	return &config, nil
}
