package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p *person) changeAge(newAge int) {
	p.age = newAge
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1.age)
	p1.changeAge(21)
	fmt.Println(p1.age)
}
