package proxy

import (
	"flag"
)

var (
	ListenAddressFlag = flag.String("listenAddress", "127.0.0.1:8040", "Address the service should listen at.")
)
