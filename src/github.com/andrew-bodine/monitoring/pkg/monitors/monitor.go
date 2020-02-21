package monitors

import (
	"flag"
)

var (
	VerboseFlag = flag.Bool("v", false, "Enables verbose logging.")
)

type Monitor interface {
	Run(stopCh <-chan struct{}) error
	GetState() interface{}
}
