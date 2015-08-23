package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)

	var c int = &a

	fmt.Println(c)
}
