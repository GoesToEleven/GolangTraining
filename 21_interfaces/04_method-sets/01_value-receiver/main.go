package main

import "fmt"

type person struct {
	first string
	last string
	age   int
}

func (p *person) fullName() string {
	fmt.Printf("Inside method:  %p\n", &p)
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := &person{"Miss", "Moneypenny", 18}

	// address in main is different from address in method
	fmt.Println(p1.fullName())
	fmt.Printf("Inside main: %p\n", &p1)

	// method works for either a value or pointer
	fmt.Println(p2.fullName())
}

// p1 is the receiver value for the call to fullName
// fullName is operating on its own value of p1
// page 117 MEAP