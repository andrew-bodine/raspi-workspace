package main

import (
	"flag"
	"fmt"

	"github.com/andrew-bodine/monitoring/pkg/monitors"
	vibrationMonitor "github.com/andrew-bodine/monitoring/pkg/vibration-monitor"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	flag.Parse()

	// Open and map memory to access gpio.
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	vibrConfig, err := vibrationMonitor.NewConfigFromFlags()
	if err != nil {
		fmt.Println("Failed to get vibration monitor config:", err.Error())
	}

	vibrStopCh := monitors.SetupSignalHandler()

	vibrMonitor, err := vibrationMonitor.NewVibrationMonitor(vibrConfig)
	if err != nil {
		fmt.Println("Failed to construct vibration monitor:", err.Error())
	}

	if err = vibrMonitor.Run(vibrStopCh); err != nil {
		fmt.Println("Failed to run vibration monitor:", err.Error())
	}
}
