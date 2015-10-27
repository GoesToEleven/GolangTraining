package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := Person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
}

/*
	p1.Greeting()
	p2.Person.Greeting()
	fmt.Println(p2.Name)
	fmt.Println(p2.Person.Name)

*/
