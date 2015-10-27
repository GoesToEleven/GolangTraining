package main

import "fmt"

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
