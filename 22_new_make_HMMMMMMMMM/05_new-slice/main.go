package main

import "fmt"

func main() {
	var members *[]string = new([]string) // new returns a pointer
	fmt.Println(members)
	fmt.Println(*members)
}
