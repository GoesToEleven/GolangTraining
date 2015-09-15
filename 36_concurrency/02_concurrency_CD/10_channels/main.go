package main

import (
	"fmt"
)

func iterate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
}

func printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	var c chan int = make(chan int)

	go iterate(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
