package example

import "testing"

type testpair struct {
	values []int
	sum    int
}

var tests = []testpair{
	{[]int{1, 2}, 3},
	{[]int{1, 1, 1, 1, 1, 1}, 6},
	{[]int{-1, 1}, 0},
}

func TestSum(t *testing.T) {
	for _, pair := range tests {
		v := Sum(pair.values...)
		if v != pair.sum {
			t.Error(
				"For", pair.values,
				"expected", pair.sum,
				"got", v,
			)
		}
	}
}
