package main

import "fmt"

type myType []int

func main() {
	var x myType = []int{32, 44, 57}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", []int(x))
}
