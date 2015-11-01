package example

// Sum finds the sum of all inputs.
func Sum(xs ...int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}
