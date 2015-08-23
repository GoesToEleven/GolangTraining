package main

import "fmt"

var a string = "this is stored in the variable a"
var b, c string = "stored in b", "stored in c"

func main() {

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c is not being used - ") // invalid code
}