package examples

import (
	"fmt"
	"time"
)

func Channels() {
	messages := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "ping"
	}()

	// Since the receiver waits for a channel to send data, there is no need for manual blocking operation
	msg := <-messages
	fmt.Println(msg)
}
