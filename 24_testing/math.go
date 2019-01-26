package math

// Adder returns the sum of two or more integers
func Adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}
