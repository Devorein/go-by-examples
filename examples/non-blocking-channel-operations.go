package examples

import "fmt"

func NonBlockChannelOperations() {
	c1 := make(chan string)
	c2 := make(chan string)

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	default:
		fmt.Println("No messages received")
	}

	select {
	case c1 <- "Message":
		fmt.Println("Sent message to channel 1")
	default:
		fmt.Println("No mesages sent")
	}
}
