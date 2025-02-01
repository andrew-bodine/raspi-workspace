package temperature_monitor

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func NewTemperatureMonitor(config interface{}) (monitors.Monitor, error) {
	temperatureMonitorConfig, ok := config.(*TemperatureMonitorConfig)
	if !ok {
		return nil, errors.New("Failed to cast config to a TemperatureMonitorConfig")
	}

	uid := uuid.NewV4()

	tm := temperatureMonitor{
		config: temperatureMonitorConfig,
		uid:    uid.String(),
		dht11:  dht11,
		state:  TemperatureMonitorStateReady,
	}

	return &tm, nil
}

type temperatureMonitor struct {
	config *TemperatureMonitorConfig

	uid   string
	dht11 *DHT11
	state TemperatureMonitorState
}

// Implement the monitors.Monitor interface.
func (tm *temperatureMonitor) Run(stopCh <-chan struct{}) error {
	tm.config.Logger.Info("Starting temperature monitor", zap.String("uid", tm.uid))
	defer tm.config.Logger.Info("Stopping temperature monitor", zap.String("uid", tm.uid))

	timer := time.NewTimer(tm.config.GPIOPinPollRate)

	var lastTempC float64

	for {
		select {
		case _, ok := <-stopCh:
			if !ok {
				return nil
			}
		case _ = <-timer.C:

			result := tm.dht11.Read()

			timer = time.NewTimer(tm.config.GPIOPinPollRate)

			if result.Error != ErrorNoError {
				tm.config.Logger.Debug("Failed to read temperature and humidity sensor",
					zap.String("error", string(result.Error)),
					zap.String("message", result.Message),
				)
				continue
			}

			currTempC = result.Temperature

			if lastTempC != currTempC {
				// Convert °C to °F.
				tm.config.Logger.Info("Current sensor readings",
					zap.String("temperature", fmt.Sprintf("%.f°F", currTempC*(9/5)+32)),
				)

				lastTempC = currTempC
			}

			break
		}
	}

	return nil
}


// Implement the monitors.Monitor interface.
func (vm *temperatureMonitor) Data() interface{} {
        return nil
}

// Implement the monitors.Monitor interface.
func (tm *temperatureMonitor) GetState() interface{} {
	return tm.state
}
