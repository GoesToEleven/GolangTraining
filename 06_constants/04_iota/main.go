package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
