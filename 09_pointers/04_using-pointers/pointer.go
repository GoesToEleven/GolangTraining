package main

import "fmt"

func main() {

	a := 43
	fmt.Println(a)

	var b *int = &a
	fmt.Println(b) // b knows the memory address where a stored 43

	*b = 42 // b says, "The value at this address, change it to 42"
	fmt.Println(a) // a looks to see what it's value is and prints 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large data structures
}
