package main

import (
	"flag"
	"log"
	"net"

	"github.com/andrew-bodine/raspi/release/pkg/os"
	pb "github.com/andrew-bodine/raspi/release/pkg/protobufs"
	"github.com/andrew-bodine/raspi/release/pkg/proxy"
)

const (
	// TODO: Make this read the containing directory to determine service name.
	serviceName = "raspi-release-proxy"
)

func main() {
	flag.Parse()

	stopCh := os.SetupSignalHandler()

	server := proxy.NewReleaseServer(stopCh, tempReleaseLister)

	listener, err := net.Listen("tcp", *proxy.ListenAddressFlag)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println(serviceName, "starting to listen at", *proxy.ListenAddressFlag)
	server.Serve(listener)
}

func tempReleaseLister(opts *proxy.ReleaseListOptions) []pb.Release {
	return []pb.Release{}
}
