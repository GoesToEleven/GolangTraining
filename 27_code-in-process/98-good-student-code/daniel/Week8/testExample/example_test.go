package example

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{5, 10}, 15},
	}

	for _, c := range tests {
		res := Sum(c.input...)
		if res != c.expected {
			t.Logf("Expected %d, but got %d", c.expected, res)
			t.Fail()
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}
