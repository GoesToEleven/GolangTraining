package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)	// 43
	fmt.Println(&a) // 0x20818a220

	var b *int = &a // valid
	fmt.Println(b) // 0x20818a220

	var c int = &a // invalid
}
