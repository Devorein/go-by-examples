package examples

import (
	"fmt"
	"time"
)

func doubleValue(workerId int, inputCh <-chan int, outputCh chan<- int) {
	for i := range inputCh {
		fmt.Println("Worker id:", workerId, "Started input", i)
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Worker id;", workerId, "Finished input", i)
		outputCh <- i * 2
	}
}

func WorkerPools() {
	const jobsNum = 5
	inputs := make(chan int, jobsNum)
	outputs := make(chan int, jobsNum)

	for id := 0; id < 3; id++ {
		go doubleValue(id, inputs, outputs)
	}

	for i := 1; i <= jobsNum; i++ {
		inputs <- i
	}
	close(inputs)

	for a := 1; a <= jobsNum; a++ {
		fmt.Println(<-outputs)
	}
}
