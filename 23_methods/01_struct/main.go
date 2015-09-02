package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) fullName() string {
	return p.fname + p.lname
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1.fname)
	fmt.Println(p1.lname)
	fmt.Println(p1.age)
	fmt.Println(p1.fullName())
}
