package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}

	fmt.Println(time.Since(start))
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
