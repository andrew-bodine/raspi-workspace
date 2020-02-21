package main

import (
	"flag"
	"sync"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
	temperatureMonitor "github.com/andrew-bodine/monitoring/pkg/temperature-monitor"
	vibrationMonitor "github.com/andrew-bodine/monitoring/pkg/vibration-monitor"
	rpio "github.com/stianeikeland/go-rpio"
	"go.uber.org/zap"
)

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if *monitors.VerboseFlag {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Open and map memory to access gpio.
	if err := rpio.Open(); err != nil {
		logger.Error("Failed to open and map memory for GPIO access:", zap.Error(err))
		return
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = runVibrationMonitor(logger); err != nil {
			logger.Error("Failed to run vibration monitor:", zap.Error(err))
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = runTemperatureMonitor(logger); err != nil {
			logger.Error("Failed to run temperature monitor:", zap.Error(err))
			return
		}
	}()

	wg.Wait()

	logger.Info("All monitors have shutdown, terminating monitoring.")
}

func runVibrationMonitor(logger *zap.Logger) error {
	config, err := vibrationMonitor.NewConfigFromFlags()
	if err != nil {
		return err
	}

	if config == nil {
		return nil
	}

	config.Logger = logger

	stopCh := monitors.SetupSignalHandler()

	monitor, err := vibrationMonitor.NewVibrationMonitor(config)
	if err != nil {
		return err
	}

	if err = monitor.Run(stopCh); err != nil {
		return err
	}

	return nil
}

func runTemperatureMonitor(logger *zap.Logger) error {
	config, err := temperatureMonitor.NewConfigFromFlags()
	if err != nil {
		return err
	}

	if config == nil {
		return nil
	}

	config.Logger = logger

	stopCh := monitors.SetupSignalHandler()

	monitor, err := temperatureMonitor.NewTemperatureMonitor(config)
	if err != nil {
		return err
	}

	if err = monitor.Run(stopCh); err != nil {
		return err
	}

	return nil
}
