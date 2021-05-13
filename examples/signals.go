package examples

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signals() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("Awaiting signal")
	<-done
	fmt.Println("Exiting")
}
