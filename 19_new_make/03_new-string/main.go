package main

import "fmt"

func main() {
	var name *string = new(string) // new returns a pointer
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Println("") // this is what *name is, an empty string
}
