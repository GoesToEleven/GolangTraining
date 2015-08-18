package main

import "fmt"

func main() {

	a := "Hello World!"
	b := false

	fmt.Println("a - ", a)
	fmt.Println("address - ", &a)
	fmt.Printf("%%v - %v\n", a)
	fmt.Printf("%%T - %T\n", a)
	fmt.Printf("%%t - %t\n", b)
	fmt.Printf("%%s - %s\n", a)
	fmt.Printf("%%q - %q\n", a)
}