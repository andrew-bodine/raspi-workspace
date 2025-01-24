package monitors

// Monitor declares the shared runtime monitor interface.
//
//go:generate counterfeiter -o fakes/fake_monitor.go --fake-name FakeMonitor . Monitor
type Monitor interface {
	Run(stopCh <-chan struct{}) error

        // Get the current data reading from the monitor.
        GetData() interface{}

	// Get the current lifecycle state of the monitor.
	GetState() interface{}
}
