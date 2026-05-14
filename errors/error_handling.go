package errorhandling

import (
	"fmt"
)

// using fmt.Errorf to wrap errors with context
func process(val int) error {
	if val < 0 {
		return fmt.Errorf("process failed: negative value %d not allowed", val)
	}
	return nil
}

func ErrorHandling() {
	fmt.Println(process(-10))
}
