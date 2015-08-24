package main

import "fmt"

func main() {

	a := "Hello World!"

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%t\n", a)
	fmt.Printf("%s\n", a)
	fmt.Printf("%q\n", a)
}