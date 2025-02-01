package registry

import (
  "github.com/andrew-bodine/raspi/monitor/pkg/monitors"
)

// Registry declares the shared runtime registry interface.
//
//go:generate counterfeiter -o fakes/fake_registry.go --fake-name FakeRegistry . Registry
type Registry interface {

    // 
    GetRegisteredMonitors() ([]monitors.Monitor, error)

    // 
    GetRegisteredMonitor() (monitors.Monitor, error)

    //
    RegisterMonitor(string, monitors.Monitor) error

    //
    UnregisterMonitor(string) error
}
