package concurrency

import (
	"fmt"
	"time"
)

func concurrencyDemoFunc(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	go concurrencyDemoFunc("goroutine")
	concurrencyDemoFunc("direct")
	
	time.Sleep(time.Millisecond * 100)
}
