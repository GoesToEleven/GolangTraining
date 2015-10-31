package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p *person) changeAge(newAge int) {
	p.age = newAge
	fmt.Printf("Inside method: %p\n", &p)
}

func main() {
	p1 := person{"James", "Bond", 20}
	p1.changeAge(21)
	fmt.Printf("Inside main: %p\n", &p1)
}

// passes the reference; the memory address
