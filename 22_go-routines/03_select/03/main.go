package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		for {
			select {
			case c1 <- "Writer #1":
			case c2 <- "Writer #2":
			}
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			// first one available will be selected
			fmt.Println("Got message 1", msg1)
		case msg2 := <-c2:
			// first one available will be selected
			fmt.Println("Got message 1", msg2)
		}
	}
	var input string
	fmt.Scanln(&input)
}
