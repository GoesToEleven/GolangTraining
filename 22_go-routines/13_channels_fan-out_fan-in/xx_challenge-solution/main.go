package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	f := fanOut(in)

	for n := range fanIn(f) {
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
	fmt.Println(len(out))
	return out
}

func fanOut(in <-chan int) []<-chan int {
	wg := sync.WaitGroup{}
	fmt.Println(len(in))
	wg.Add(len(in))

	xc := make([]<-chan int, 0)
	out := make(chan chan int)

	go func() {
		for n := range in {
			c := make(chan int)
			c <- factorial(n)
			close(c)
			out <- c
		}
		close(out)
	}()

	go func() {
		for n := range out {
			xc = append(xc, n)
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Printf("%T\n", xc)
	fmt.Printf("%v\n", xc)
	return xc
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func fanIn(xc []<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(xc))

	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range xc {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
