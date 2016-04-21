package example

import "testing"

func TestSum(t *testing.T) {
	var n int
	n = Sum(1, 2)
	if n != 3 {
		t.Error("Expected 3, got ", n)
	}
}
