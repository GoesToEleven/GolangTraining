package main

import (
	"fmt"
	"strconv"
)

func main() {

	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	x := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y := strconv.FormatInt(-42, 16)
	z := strconv.FormatUint(42, 16)

	fmt.Println(w, x, y, z)
}
