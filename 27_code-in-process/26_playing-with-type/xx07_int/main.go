package main

import "fmt"

type myType int

func main() {
	var x myType = 32
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", int(x))
}
