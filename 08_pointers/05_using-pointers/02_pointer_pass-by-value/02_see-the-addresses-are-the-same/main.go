package main

import "fmt"

func zero(x *int) {
	fmt.Println(x)
	*x = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
