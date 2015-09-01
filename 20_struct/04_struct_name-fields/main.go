package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{name: "James", age: 20}
	fmt.Println(p1.name)
}
