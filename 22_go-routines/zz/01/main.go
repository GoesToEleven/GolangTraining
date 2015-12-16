package main

import "fmt"

var c = make(chan int)

func main() {
	go incrementor()
	puller()
	fmt.Println("Program Exiting")
}

func incrementor() {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func puller() {
	for n := range c {
		fmt.Println(n)
	}
}

// go run -race main.go
// vs
// go run main.go
