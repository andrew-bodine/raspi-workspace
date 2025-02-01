package monitors

// Monitor declares the shared runtime monitor interface.
//
//go:generate counterfeiter -o fakes/fake_monitor.go --fake-name FakeMonitor . Monitor
type Monitor interface {

	// Instructs the monitor process to start running and provides a channel
	// for knowing when it's creator wants it to stop.
	Run(stopCh <-chan struct{}) error

        // Get the current data reading from the monitor.
        Data() interface{}

	// Get the current lifecycle state of the monitor.
	GetState() interface{}
}
