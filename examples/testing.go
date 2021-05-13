package examples

import (
	"fmt"
	"testing"
)

func IntMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestIntMinBasic(t *testing.T) {
	res := IntMin(2, 3)
	if res != 2 {
		t.Errorf("IntMain(%d, %d) = %d, want %d", 2, 3, res, 2)
	}
}

func TestIntMinTable(t *testing.T) {
	samples := []struct {
		a    int
		b    int
		want int
	}{
		{
			1, 2, 1,
		},
		{
			3, 2, 2,
		},
		{
			-3, 2, -3,
		},
	}

	for _, s := range samples {
		testname := fmt.Sprintf("IntMin(%d,%d) = %d", s.a, s.b, s.want)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(s.a, s.b)
			if ans != s.want {
				t.Errorf("got %d, want %d", ans, s.want)
			}
		})
	}
}
