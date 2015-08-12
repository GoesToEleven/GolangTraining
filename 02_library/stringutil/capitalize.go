package stringutil

func reverseTwo (s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// the name of this file doesn't correspond to what the code does
// this would not be good programming to do it this way
// this is just to demonstrate how the reverseTwo function can be in this file
// in the stringutil package
// and be used by the reverse.go file
// also in the same stringutil package
