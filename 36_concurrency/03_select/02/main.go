package main

import (
	"fmt"
	"math/rand"
	"time"
)

func writer(msg string, c chan string) {
	for {
		time.Sleep(time.Second + time.Duration(rand.Intn(100))*time.Millisecond)
		c <- msg
	}
}

func main() {
	c1, c2 := make(chan string), make(chan string)

	go writer("Writer #1", c1)
	go writer("Writer #2", c2)

	for {
		select {
		case msg1 := <-c1:
			// first one available will be selected
			fmt.Println("Got message 1", msg1)
		case msg2 := <-c2:
			// first one available will be selected
			fmt.Println("Got message 1", msg2)
		case <-time.After(time.Second):
			fmt.Println("TIMEOUT")
		default:
			fmt.Println("Default")
		}
	}
	var input string
	fmt.Scanln(&input)
}
