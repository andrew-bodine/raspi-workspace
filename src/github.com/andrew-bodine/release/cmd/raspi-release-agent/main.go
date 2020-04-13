package main

import (
	"fmt"
	"log"

	"github.com/andrew-bodine/release/pkg/os"
	pb "github.com/andrew-bodine/release/pkg/protobufs"
)

const (
	SystemdInstallPath = "/usr/lib/systemd/system"

	RaspiReleaseInstallPath = "/opt/raspi"
	RaspiReleaseBinPath     = RaspiReleaseInstallPath + "/bin"
	RaspiReleasesPath       = RaspiReleaseInstallPath + "/releases"
)

var releaseQueue chan *pb.Release

func main() {
	stopCh := os.SetupSignalHandler()

	releaseQueue = make(chan *pb.Release)

	// Start a goroutine to sync local release cache with the release proxy.
	go runReleaseCacheSync(stopCh, releaseQueue)

	runReleaseReconciler(stopCh, releaseQueue)
}

func runReleaseCacheSync(stopCh <-chan struct{}, releaseQueue chan *pb.Release) {
	fmt.Println("starting agent release cacher")
	defer fmt.Println("terminating agent release cacher")

	for {
		select {
		case <-stopCh:
			return
		}
	}
}

func runReleaseReconciler(stopCh <-chan struct{}, releaseQueue chan *pb.Release) {
	fmt.Println("starting agent release reconciler")
	defer fmt.Println("terminating agent release reconciler")

	for {
		select {
		case <-stopCh:
			return
		case release := <-releaseQueue:
			log.Println("RECONCILER: was notified of new release on queue.")
			log.Println(release)
		}
	}
}
