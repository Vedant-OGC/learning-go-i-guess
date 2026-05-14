package errorhandling

import (
	"errors"
	"fmt"
)

func errorDemoFunc(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func ErrorHandling() {
	res, err := errorDemoFunc(42)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", res)
	}
}
