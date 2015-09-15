package main

import (
	"fmt"
)

func helloWorld(i int) {
	fmt.Println(i, "Hello world")
}

func main() {
	for i := 0; i < 5; i++ {
		go helloWorld(i)
	}
	var input string
	fmt.Scanln(&input)
}
