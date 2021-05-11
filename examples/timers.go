package examples

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 completed")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 completed")
	}()

	timer2.Stop()
	fmt.Println("Timer 2 stopped")
	time.Sleep(2 * time.Second)
}
