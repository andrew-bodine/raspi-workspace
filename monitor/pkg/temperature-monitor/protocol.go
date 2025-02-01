package temperature_monitor

import (
	"fmt"

	"github.com/andrew-bodine/raspi/monitor/pkg/monitors"
)

type DHT11Result struct {
	Temperature float32
}

func (dht11 *DHT11) Read() *DHT11Result {
	return nil
}

