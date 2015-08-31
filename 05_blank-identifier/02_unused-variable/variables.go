package main

import "fmt"

var a string = "this is stored in the variable a"
var b, _ string = "stored in b", "stored then thrown away"

func main() {

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c is not being used -  and this is no problem")
}