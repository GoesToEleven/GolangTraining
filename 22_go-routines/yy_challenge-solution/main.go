package main

import (
	"fmt"
)

var c = make(chan int)

func main() {
	for i := 0; 1 < 10; i++ {
		for j := 4; i < 14; j++ {
			go func(){
				c <- factorial(j)
			}()
		}
	}
	close(c)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) int{
	total := 1
	for i := n; i > 0; i-- {
			total *= i
	}
	return total
}
