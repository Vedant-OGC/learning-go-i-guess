package errorhandling

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func CustomErrors() {
	err := &MyError{Code: 500, Message: "Internal Fail"}
	
	// errors.As checks if an error matches target type
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Println("matched error code:", myErr.Code)
	}
}
