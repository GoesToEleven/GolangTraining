package main

import "fmt"

func main() {
	var member *bool = new(bool) // new returns a pointer
	fmt.Println(member)
	fmt.Println(*member)
}
