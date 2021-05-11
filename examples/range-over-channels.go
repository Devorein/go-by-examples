package examples

import "fmt"

func RangeOverChannels() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	for elem := range ch {
		fmt.Println(elem)
	}
}
