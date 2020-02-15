package vibration_monitor

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
	uuid "github.com/satori/go.uuid"
	rpio "github.com/stianeikeland/go-rpio"
)

var (
	VibrPinFlag         = flag.Int("vibrPin", -1, "Raspberry Pi GPIO pin to monitor for vibrations.")
	VibrPinPollRateFlag = flag.Duration("vibrPinPollRate", time.Millisecond*500, "How often to check the state of --vibrPin.")
	VibrMinPeriodFlag   = flag.Duration("vibrMinPeriod", time.Second*15, "Minimum period of sustained vibration that constitutes an event.")
)

func NewConfigFromFlags() (*VibrationMonitorConfig, error) {
	if *VibrPinFlag == -1 {
		return nil, errors.New("Required flag --vibrPin missing, please specify and try again.")
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
	log("INFO", "Starting vibration monitor", vm.uid)
	defer log("INFO", "Stopping vibration monitor", vm.uid)

	for {
		select {
		case _, ok := <-stopCh:
			if !ok {
				return nil
			}
		default:
			break
		}

		if vm.config.GPIOPin.Read() == rpio.High {
			log("DEBUG", "Pin is reading as High.")

			switch vm.state {
			case VibrationMonitorStateStill:
				vm.state = VibrationMonitorStateHalfVibrating
				vm.vibrationPeriodTimer = time.NewTimer(vm.config.MinVibrationPeriod)
				break
			case VibrationMonitorStateHalfVibrating:
				select {
				case _ = <-vm.vibrationPeriodTimer.C:
					vm.state = VibrationMonitorStateVibrating
					log("INFO", "Vibration period begins after initial detection grace period", vm.config.MinVibrationPeriod)
				default:
					break
				}
				break
			}
		} else {
			log("DEBUG", "Pin is reading as Low.")

			switch vm.state {
			case VibrationMonitorStateVibrating:
				log("INFO", "Vibration period ends")
			}

			vm.state = VibrationMonitorStateStill
		}

		time.Sleep(vm.config.GPIOPinPollRate)
	}

	return nil
}

// Implement the monitors.Monitor interface.
func (vm *vibrationMonitor) GetState() interface{} {
	return vm.state
}

func log(level string, msgs ...interface{}) {
	now := time.Now()
	fmt.Println(level, now.String(), msgs)
}
