package main

import "fmt"

func main() {
	x := 42
	another()
}

func another() {
	fmt.Println(x)
}