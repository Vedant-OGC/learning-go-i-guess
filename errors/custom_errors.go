package errorhandling

import "fmt"

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func runCustomErrorDemo() error {
	return &MyError{Code: 404, Message: "Not Found"}
}

func CustomErrors() {
	err := runCustomErrorDemo()
	if err != nil {
		fmt.Println(err)
	}
}
