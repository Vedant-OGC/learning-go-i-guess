package errorhandling

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	return nil
}

func ErrorHandling() {
	err := checkAge(-5)
	if err != nil {
		// handling error
		fmt.Println("caught:", err)
	}
}
