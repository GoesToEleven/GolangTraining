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
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
