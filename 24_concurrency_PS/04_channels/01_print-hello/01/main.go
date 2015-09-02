package main

import "fmt"

func main() { 
	ch := make(chan string)
	fmt.Println(<-ch) 	// program waits here to drain the channel
						// since there is nothing to receive from the channel
						// our program is in a deadlock
}