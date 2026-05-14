package errorhandling

import "fmt"

func PanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	
	panic("something bad happened")
}
