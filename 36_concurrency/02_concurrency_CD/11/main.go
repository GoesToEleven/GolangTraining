package main

import "fmt"

func pinger(c chan string) {
	c <- "ping"
}

func ponger(c1 chan string, c2 chan string) {
	msg := <-c1
	fmt.Println("PONGER GOT", msg)
	c2 <- "pong"
}

func printer(c2 chan string) {
	fmt.Println(<-c2)
}

func main() {
	c1, c2 := make(chan string), make(chan string)

	go pinger(c1)
	go ponger(c1, c2)
	go printer(c2)

	var input string
	fmt.Scanln(&input)
}


