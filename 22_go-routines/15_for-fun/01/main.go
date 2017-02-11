package main

import (
	"fmt"
)

func main() {
	m := map[int]int{}
	m[4] = 7
	m[3] = 87
	m[72] = 19

	ch := make(chan int)

	// for closing ch
	ch2 := make(chan int)
	go func() {
		var i int
		for n := range ch2 {
			i += n
			if i == len(m) {
				close(ch)
			}
		}
	}()

	go func() {
		for _, v := range m {
			ch <- v
			ch2 <- 1
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}

	// good housekeeping
	close(ch2)
}