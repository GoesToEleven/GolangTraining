package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c

		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)

	go ping(c)
	go pong(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
