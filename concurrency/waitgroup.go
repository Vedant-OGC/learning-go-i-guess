package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println("worker complete:", id)
		}(i)
	}
	wg.Wait()
}
