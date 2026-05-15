package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Goroutines() {
	var counter int64
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1) // atomic operation
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter (atomic, safe):", counter)
}
