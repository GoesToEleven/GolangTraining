package main

import (
	"strings"
	"fmt"
)

func main() {
	phrase := "These are the times that try men's souls\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)
	// closing a channel only closes the ability to send onto the channel
	// data on the channel remains on channel
	// and channel can still be received from
	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

	// you can't send on a closed channel:
	ch <- "test"

}

