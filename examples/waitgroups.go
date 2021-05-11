package examples

import (
	"fmt"
	"sync"
	"time"
)

func waitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("starting worker %d\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go waitGroup(i, &wg)
	}

	wg.Wait()
	fmt.Println("Finished all workers")
}
