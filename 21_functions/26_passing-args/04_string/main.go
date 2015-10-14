package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}

func changeMe(x string) {
	fmt.Println(x) // Todd
	x = "Apollo"
	fmt.Println(x) // Apollo
}
