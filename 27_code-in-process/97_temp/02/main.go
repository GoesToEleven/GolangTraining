package main

import (
	"fmt"
)

func init() {
	fmt.Println("Who ran first?", x)
}

func main() {
	fmt.Println("Hello world.")
}

var x int = 17
