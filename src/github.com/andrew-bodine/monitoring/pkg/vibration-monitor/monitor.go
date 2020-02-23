package vibration_monitor

import (
	"errors"
	"time"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
	uuid "github.com/satori/go.uuid"
	rpio "github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"
)

func NewVibrationMonitor(config interface{}) (monitors.Monitor, error) {
	vibrationMonitorConfig, ok := config.(*VibrationMonitorConfig)
	if !ok {
		return nil, errors.New("Failed to cast config to a VibrationMonitorConfig")
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	vm := vibrationMonitor{
		config: vibrationMonitorConfig,
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
