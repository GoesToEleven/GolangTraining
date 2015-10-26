package main

import "fmt"

func main() {
	var age *map[int]int = new(map[int]int) // new returns a pointer
	fmt.Println(age)
	fmt.Println(*age)

	var bday map[int]int = make(map[int]int) // make is only for slice, map, channel
	fmt.Println(bday)
	fmt.Println(bday[0])
	fmt.Println(bday[10])
	fmt.Println(bday[1000])
	fmt.Println(bday[9999999999999])
}
