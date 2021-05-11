package examples

import (
	"fmt"
	"time"
)

func MultiChannelSync() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "World"
		c1 <- "Again"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
