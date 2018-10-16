package main

import "fmt"


// var keyword
var y = 9.99

func main() {
	// Short delcation operator
	x := 44.5

	fmt.Println(x)

	foo()
}

func foo() {
	fmt.Println("in foo",y)
}

