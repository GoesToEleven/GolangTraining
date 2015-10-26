package main

import "fmt"

func main() {
	var members *[]string = new([]string) // new returns a pointer
	fmt.Println(members)
	fmt.Println(*members)

	var staff []string = make([]string, 40, 100) // make is only for slice, map, channel
	fmt.Println(staff)
}
