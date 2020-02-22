package temperature_monitor

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
	uuid "github.com/satori/go.uuid"
	rpio "github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"
)

var (
	TempPinFlag         = flag.Int("tempPin", -1, "Raspberry Pi GPIO pin to monitor for temperature and humidity changes.")
	TempPinPollRateFlag = flag.Duration("tempPinPollRate", time.Second*5, "How often to check the state of --tempPin.")
)

func NewConfigFromFlags() (*TemperatureMonitorConfig, error) {
	if *TempPinFlag == -1 {
		return nil, nil
	}

	pin := rpio.Pin(*TempPinFlag)

	config := TemperatureMonitorConfig{
		GPIOPin:         pin,
		GPIOPinPollRate: *TempPinPollRateFlag,
	}

	return &config, nil
}

type TemperatureMonitorConfig struct {
	Logger *zap.Logger

	// Represents the GPIO pin that this monitor is pay attention to.
	GPIOPin monitors.GoRaspberryPiIOPin

	GPIOPinPollRate time.Duration
}

type TemperatureMonitorState string

const (
	TemperatureMonitorStateReady = "Ready"
)

func NewTemperatureMonitor(config *TemperatureMonitorConfig) (monitors.Monitor, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	dht11 := &DHT11{Pin: config.GPIOPin}

	tm := temperatureMonitor{
		config: config,
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

	var lastTempC, lastHumidity float64

	for {
		select {
		case _, ok := <-stopCh:
			if !ok {
				return nil
			}
		case _ = <-timer.C:

			// result := tm.dht11.ReadCurrentTempAndHumidity()
			result := tm.dht11.Read()

			timer = time.NewTimer(tm.config.GPIOPinPollRate)

			if result.Error != ErrorNoError {
				tm.config.Logger.Debug("Failed to read temperature and humidity sensor",
					zap.String("error", string(result.Error)),
					zap.String("message", result.Message),
				)
				continue
			}

			currTempC, tempConvErr := strconv.ParseFloat(fmt.Sprintf("%d.%d", int(result.Temperature[0]), int(result.Temperature[1])), 32)
			currHumidity, humConvErr := strconv.ParseFloat(fmt.Sprintf("%d.%d", int(result.Humidity[0]), int(result.Humidity[1])), 32)

			if tempConvErr != nil || humConvErr != nil {
				tm.config.Logger.Debug("Failed to parse floats from sensor readings",
					zap.Error(tempConvErr),
					zap.Error(humConvErr),
				)
				continue
			}

			if lastTempC != currTempC || lastHumidity != currHumidity {
				// Convert °C to °F.
				tm.config.Logger.Info("Current sensor readings",
					zap.String("temperature", fmt.Sprintf("%.f°F", currTempC*(9/5)+32)),
					zap.String("humidity", fmt.Sprintf("%d.%d%%", int(result.Humidity[0]), int(result.Humidity[1]))),
				)

				lastTempC = currTempC
				lastHumidity = currHumidity
			}

			break
		}
	}

	return nil
}

// Implement the monitors.Monitor interface.
func (tm *temperatureMonitor) GetState() interface{} {
	return tm.state
}
