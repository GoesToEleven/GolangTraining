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

	close(ch) // closing the channel removes deadlock

	for {
		if msg, ok := <-ch; ok { // when channel is closed
			fmt.Print(msg + " ") // this for loop will no longer be waiting
		} else { // to receive something from channel
			break // the loop will break
		}
	}

}
