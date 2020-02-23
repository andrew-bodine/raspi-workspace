package monitors

import (
	"fmt"
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
		fmt.Println("Received stop signal, attempting graceful termination...")
		close(stop)

		<-c
		fmt.Println("Received stop signal, terminating immediately!")
		os.Exit(1)
	}()

	return stop
}
