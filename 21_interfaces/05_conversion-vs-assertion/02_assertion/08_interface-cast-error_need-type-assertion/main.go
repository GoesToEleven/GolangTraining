package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", int(val))
	//	fmt.Printf("%T\n", val.(int))
}
