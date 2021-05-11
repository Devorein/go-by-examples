package examples

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ChannelSync() {
	done := make(chan bool)
	go worker(done)
	fmt.Println(<-done)
	fmt.Println("Last")
}
