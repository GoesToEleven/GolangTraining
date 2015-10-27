package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func main() {
	c1 := Customer{"Todd", 44}
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x8201e4120
}

func changeMe(z *Customer) {
	fmt.Println(z)       // &{Todd 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x8201e4120

}
