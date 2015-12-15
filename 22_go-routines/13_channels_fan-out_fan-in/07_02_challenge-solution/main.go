package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	f := fanOut(in)

	for n := range f {
		fmt.Println(n)
		fmt.Printf("%T", n)
	}

	fmt.Printf("%T\n", f)

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

func fanOut(in <-chan int) <-chan <-chan int {

	wg := sync.WaitGroup{}
	out := make(chan (<-chan int))
	defer close(out)

	for n := range in {
		go func() {
			out <- factorial(n)
		}()
	}

	return out
}

func factorial(n int) <-chan int {
	c := make(chan int)
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	c <- total
	close(c)
	return c
}
