package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])
}
