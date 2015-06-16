package recursion

import (
	"testing"
)

func TestFib(t *testing.T) {
	n := Fib(20)
	if n != 6765 {
		t.Error("Expected Fib(20) to equal 6765, not %v", n)
	}
}
