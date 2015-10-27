package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p1 person
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Printf("%T\n", p1)
}

// always use var to create and
// initialize a variable to its zero value
