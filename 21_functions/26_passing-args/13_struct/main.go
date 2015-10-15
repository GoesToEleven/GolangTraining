package main

import "fmt"

type Customer struct {
	name string
	age int
}

func main() {
	c1 := Customer{"Todd", 44}
	fmt.Println(&c1) // &{Todd 44}

	changeMe(c1)

	fmt.Println(c1) //{Todd 44}
	fmt.Println(&c1) //&{Todd 44}
}

func changeMe(x Customer) {
	fmt.Println(x) //{Todd 44}
	fmt.Println(&x) //&{Todd 44}
	x.name = "Apollo"
	fmt.Println(x) //{Apollo 44}
	fmt.Println(&x) //&{Apollo 44}

}
