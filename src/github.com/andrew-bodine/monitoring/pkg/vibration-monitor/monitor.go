package vibration_monitor

import (
	"flag"
	"time"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
	uuid "github.com/satori/go.uuid"
	rpio "github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"
)

var (
	VibrPinFlag         = flag.Int("vibrPin", -1, "Raspberry Pi GPIO pin to monitor for vibrations.")
	VibrPinPollRateFlag = flag.Duration("vibrPinPollRate", time.Millisecond*500, "How often to check the state of --vibrPin.")
	VibrMinPeriodFlag   = flag.Duration("vibrMinPeriod", time.Second*15, "Minimum period of sustained vibration that constitutes an event.")
)

func NewConfigFromFlags() (*VibrationMonitorConfig, error) {
	if *VibrPinFlag == -1 {
		return nil, nil
		// return nil, errors.New("Required flag --vibrPin missing, please specify and try again.")
	}

	// Get a reference to the desired pin.
	pin := rpio.Pin(*VibrPinFlag)

	// Set pin in input mode.
	pin.Input()

	config := VibrationMonitorConfig{
		GPIOPin:            pin,
		GPIOPinPollRate:    *VibrPinPollRateFlag,
		MinVibrationPeriod: *VibrMinPeriodFlag,
	}

	return &config, nil
}

// Collection of configurable settings on that are common amongst vibration
// monitor implementations.
type VibrationMonitorConfig struct {
	Logger *zap.Logger

	// Represents the GPIO pin that this monitor is pay attention to.
	GPIOPin monitors.GoRaspberryPiIOPin

	GPIOPinPollRate time.Duration

	MinVibrationPeriod time.Duration
}

type VibrationMonitorState string

const (
	VibrationMonitorStateStill         VibrationMonitorState = "Still"
	VibrationMonitorStateHalfVibrating VibrationMonitorState = "HalfVibrating"
	VibrationMonitorStateVibrating     VibrationMonitorState = "Vibrating"
)

func NewVibrationMonitor(config *VibrationMonitorConfig) (monitors.Monitor, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	vm := vibrationMonitor{
		config: config,
		uid:    uid.String(),
		state:  VibrationMonitorStateStill,
	}

	return &vm, nil
}

// Default implementation of the monitors.Monitor interface.
type vibrationMonitor struct {
	config *VibrationMonitorConfig

	uid                  string
	vibrationPeriodTimer *time.Timer
	state                VibrationMonitorState
}

// Implement the monitors.Monitor interface.
func (vm *vibrationMonitor) Run(stopCh <-chan struct{}) error {
	vm.config.Logger.Info("Starting vibration monitor", zap.String("uid", vm.uid))
	defer vm.config.Logger.Info("Stopping vibration monitor", zap.String("uid", vm.uid))

	timer := time.NewTimer(vm.config.GPIOPinPollRate)

	for {
		select {
		case _, ok := <-stopCh:
			if !ok {
				return nil
			}
		case _ = <-timer.C:
			break
		}

		if vm.config.GPIOPin.Read() == rpio.High {
			vm.config.Logger.Debug("Pin is reading as High.")

			switch vm.state {
			case VibrationMonitorStateStill:
				vm.state = VibrationMonitorStateHalfVibrating
				vm.vibrationPeriodTimer = time.NewTimer(vm.config.MinVibrationPeriod)
				break
			case VibrationMonitorStateHalfVibrating:
				select {
				case _ = <-vm.vibrationPeriodTimer.C:
					vm.state = VibrationMonitorStateVibrating
					vm.config.Logger.Info("Vibration period begins after initial detection grace period",
						zap.Duration("minVibrationPeriod", vm.config.MinVibrationPeriod),
					)
				default:
					break
				}
				break
			}
		} else {
			vm.config.Logger.Debug("Pin is reading as Low.")

			switch vm.state {
			case VibrationMonitorStateVibrating:
				vm.config.Logger.Info("Vibration period ends")
			}

			vm.state = VibrationMonitorStateStill
		}

		timer = time.NewTimer(vm.config.GPIOPinPollRate)
	}

	return nil
}

// Implement the monitors.Monitor interface.
func (vm *vibrationMonitor) GetState() interface{} {
	return vm.state
}
