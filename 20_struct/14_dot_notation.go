package main

import "fmt"

// ***********************************************
// CAPITALIZATION MATTERS
// -- capitalized: available (public)
// -- lower case: not available (private)
// ***********************************************

type Contact struct {
	name     string
	greeting string
}

func main() {

	var c = Contact{}
	c.name = "Marcus"
	c.greeting = "Hello!"

	fmt.Println(c.name)
	fmt.Println(c.greeting)

}

// we can use a STRUCT like a class

// ***********************************************
// we can also add functions (methods) to user defined types
// ***********************************************

// go is not an OOP language
// the type sustem that exists in go
// -- makes it so you don't need classes
// -- gives you more flexibility b/c you're not constrained by class requirements
// instead of using classes, we have user defined types
