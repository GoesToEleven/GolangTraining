package main

import "fmt"

func main() {
	x := "true"
	y := x == "true"
	fmt.Printf("%T, %v\n", y, y)
}
