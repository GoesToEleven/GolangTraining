package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "James"}
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
