package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's souls\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for {
		if msg, ok := <-ch; ok { // we check to see if channel is closed
			fmt.Print(msg + " ")
		} else {
			break
		}
	}
	// but we haven't closed the channel yet
	// so the for loop on line 19
	// loops through all of the words on the channel
	// then waits for another word to be put on the channel
	// and as no word is ever going to be put on the channel
	// program is in deadlock
}
