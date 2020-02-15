package monitors

type Monitor interface {
	Run(stopCh <-chan struct{}) error
	GetState() interface{}
}
