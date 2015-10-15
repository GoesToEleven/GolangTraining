package main

import "fmt"

type Customer struct {
	name string
	age int
}

func main() {
	c1 := Customer{"Todd", 44}
	fmt.Println(&c1.name) //0x8201e4120

	changeMe(c1)

	fmt.Println(c1) //{Todd 44}
	fmt.Println(&c1.name) //0x8201e4120
}

func changeMe(x Customer) {
	fmt.Println(x) //{Todd 44}
	fmt.Println(&x.name) //0x8201e4140
	x.name = "Apollo"
	fmt.Println(x) //{Apollo 44}
	fmt.Println(&x.name) //0x8201e4140

}
