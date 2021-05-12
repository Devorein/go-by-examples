package examples

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounters() {
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for i := 1; i <= 1000; i++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter", counter)
}
