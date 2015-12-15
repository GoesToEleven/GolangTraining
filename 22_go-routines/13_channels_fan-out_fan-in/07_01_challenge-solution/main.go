package main

import (
	"fmt"
)

func main() {

	in := gen()

	for n := range in {
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
