package main

import (
	"flag"
	"sync"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
	temperatureMonitor "github.com/andrew-bodine/raspi/monitor/pkg/temperature-monitor"
	vibrationMonitor "github.com/andrew-bodine/raspi/monitor/pkg/vibration-monitor"
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

		err := monitors.BuildAndRunMonitor(&monitors.ConfigAndBuilderWrapper{
			Logger:            logger,
			ConfigConstructor: vibrationMonitor.NewConfigFromFlagsWithLogger,
			MonitorBuilder:    vibrationMonitor.NewVibrationMonitor,
		})
		if err != nil {
			logger.Error("Failed to run vibration monitor:", zap.Error(err))
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := monitors.BuildAndRunMonitor(&monitors.ConfigAndBuilderWrapper{
			Logger:            logger,
			ConfigConstructor: temperatureMonitor.NewConfigFromFlagsWithLogger,
			MonitorBuilder:    temperatureMonitor.NewTemperatureMonitor,
		})
		if err != nil {
			logger.Error("Failed to run temperature monitor:", zap.Error(err))
			return
		}
	}()

	wg.Wait()

	logger.Info("All monitors have shutdown, terminating monitoring.")
}
