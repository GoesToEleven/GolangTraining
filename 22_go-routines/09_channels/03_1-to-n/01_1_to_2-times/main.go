package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done

}
