package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "James"}
	fmt.Printf("%T\n", p1)
}
