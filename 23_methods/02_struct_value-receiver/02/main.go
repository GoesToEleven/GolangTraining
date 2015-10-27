package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) fullName() string {
	fmt.Printf("Inside method:  %p\n", &p)
	return p.fname + p.lname
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1.fullName())
	fmt.Printf("Inside main: %p\n", &p1)
}

// p1 is the receiver value for the call to fullName
// fullName is operating on a copy of p1
