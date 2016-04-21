package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "James",
		age:  20,
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Printf("%T\n", p1)
}

// always use shorthand notation to
// create and initialize a variable to values
