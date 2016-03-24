package demo_test

import (
	"demo"
	"runtime"
	"testing"
)

func TestSumStrings(t *testing.T) {
	sum := demo.SumStrings("1", "2", "3")
	if sum != 6 {
		runtime.Breakpoint()
		t.Errorf("expecting sum to be 6 was %d", sum)
	}
}
