package mytests

import (
	"testing"
	"github.com/Vedant-OGC/learning-go-i-guess/packages"
)

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
	}
	
	for _, tt := range tests {
		t.Run("testing", func(t *testing.T) {
			ans := packages.Add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, ans, tt.want)
			}
		})
	}
}
