package concurrency

import (
	"fmt"
	"time"
)

func SelectStmt() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // timeout!
		fmt.Println("timeout 1")
	}
}
