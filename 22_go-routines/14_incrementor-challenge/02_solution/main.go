package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := incrementor(2)

	var count int
	for n := range c {
		count++
		fmt.Println(n + strconv.Itoa(count))
	}

	fmt.Println("Final Count:", strconv.Itoa(count))
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", k, "\tTotal Count: ")
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}
