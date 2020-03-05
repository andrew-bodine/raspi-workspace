package main

import (
	"flag"
	"log"
	"net"

	"github.com/andrew-bodine/release/pkg/os"
	"github.com/andrew-bodine/release/pkg/proxy"
)

const (
	// TODO: Make this read the containing directory to determine service name.
	serviceName = "raspi-release-proxy"
)

func main() {
	flag.Parse()

	stopCh := os.SetupSignalHandler()

	server := proxy.NewReleaseServer(stopCh)

	listener, err := net.Listen("tcp", *proxy.ListenAddressFlag)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println(serviceName, "starting to listen at", *proxy.ListenAddressFlag)
	server.Serve(listener)
}
