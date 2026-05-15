package concurrency

import (
	"fmt"
	"sync"
)

func Goroutines() {
	// race condition!
	var counter int
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // concurrent write, race condition!
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter (racy):", counter)
}
