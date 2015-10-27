package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
