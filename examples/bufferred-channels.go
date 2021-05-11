package examples

import (
	"fmt"
)

func BufferedChannels() {
	fmt.Println()

	// Without bufferring there would be no concurrent sender to the channel thus the receiver would not receive any value
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "World"

	msg1 := <-messages
	msg2 := <-messages

	fmt.Println(msg1)
	fmt.Println(msg2)
}
