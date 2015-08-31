package main

import "fmt"

func main() {
	var age *[]int = new([]int) // new returns a pointer
	fmt.Println(age)
	fmt.Println(*age)

	var bday []int = make([]int, 10, 100) // make is only for slice, map, channel
	fmt.Println(bday)
}
