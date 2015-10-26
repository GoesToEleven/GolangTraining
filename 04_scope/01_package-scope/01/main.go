package main

import "fmt"

var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
