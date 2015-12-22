package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

/*
Even easier is just to create a new variable,
using a declaration style that may seem odd but works fine in Go.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
