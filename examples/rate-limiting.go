package examples

import (
	"fmt"
	"time"
)

func RateLimiting() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	ticker := time.NewTicker(200 * time.Millisecond)

	for req := range requests {
		<-ticker.C
		fmt.Println(req)
	}

	burstLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for r := range burstRequests {
		<-burstLimiter
		fmt.Println("request", r, time.Now())
	}
}
