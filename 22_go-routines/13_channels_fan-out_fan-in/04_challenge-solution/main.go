package main

import (
	"fmt"
	"google.golang.org/appengine/internal/search"
)

func main() {

	in := gen()

	// Distribute work across goroutines that read from in.
	f := fanOut(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	for n := range in {
		go func(){

		}()
		close(out)
	}
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}