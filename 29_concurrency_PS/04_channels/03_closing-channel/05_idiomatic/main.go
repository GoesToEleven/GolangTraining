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
	
	for msg := range ch {
			fmt.Print(msg + " ")
	}
	// range knows to stop looping
	// when there is nothing left in a closed channel
}