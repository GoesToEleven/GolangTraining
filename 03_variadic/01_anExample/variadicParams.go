package main

import "fmt"

func main() {
	Greeting("nobody")
	Greeting("hello:", "Joe", "Anna", "Eileen")
}

func Greeting(prefix string, people ...string) {
	fmt.Print(prefix, " ")
	for _, person := range people {
		fmt.Print(person, ", ")
	}
	fmt.Println()
}