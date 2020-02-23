package monitors

import (
	"flag"
)

var (
	VerboseFlag = flag.Bool("v", false, "Enables verbose logging.")
)

// Monitor declares the shared runtime monitor interface.
//
//go:generate counterfeiter -o fakes/fake_monitor.go --fake-name FakeMonitor . Monitor
type Monitor interface {
	Run(stopCh <-chan struct{}) error
	GetState() interface{}
}
