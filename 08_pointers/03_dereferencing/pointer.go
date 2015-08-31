package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a) 	// 43
	fmt.Println(&a) // 0x20818a220

	var b *int = &a
	fmt.Println(b)	// 0x20818a220
	fmt.Println(*b)	// 43

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
}
