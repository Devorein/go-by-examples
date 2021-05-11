package examples

import (
	"fmt"
	"time"
)

func Timeouts() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Message 1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second):
		fmt.Println("1 sec timeout")
	}

	go func() {
		fmt.Println(time.Second)
		c2 <- "Message 2"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("3 sec timeout")
	}
}
