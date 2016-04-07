package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func(x int) {
			for q := range c {
				fmt.Println("i", x, "q", q)
			}
			done <- true
		}(i)
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
