package concurrency

import "fmt"

func Channels() {
	// producer consumer demo
	jobs := make(chan int, 5)
	done := make(chan bool)
	
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	<-done
}
