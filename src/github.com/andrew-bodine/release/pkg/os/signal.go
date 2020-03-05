package os

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SetupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	c := make(chan os.Signal, 2)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received stop signal, attempting graceful termination...")
		close(stop)
		<-c
		log.Println("Received stop signal, terminating immediately!")
		os.Exit(1)
	}()

	return stop
}
