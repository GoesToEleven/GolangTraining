package main

import "fmt"

func main() {
	var members *map[int]string = new(map[int]string) // new returns a pointer
	fmt.Println(members)
	fmt.Println(*members)
}
