package mytests

import (
	"testing"
	"github.com/Vedant-OGC/learning-go-i-guess/packages"
)

func TestAdd(t *testing.T) {
	ans := packages.Add(2, 3)
	if ans != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", ans)
	}
}
