package main

import "fmt"

func main() {

	m := 'm' // single quotes
	n := "n" // double quotes
	o := `o` // back ticks

	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)

	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", o)
}