package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	ch <- "Hello"	// program waits here until channel is drained
					// since our main thread is waiting on the above line
					// the main thread never gets to the statement below that drains the channel
					// our program is in a deadlock
	fmt.Println(<-ch)
}

/*
we can give our channel
the capacity to store messages
which will allow the sender & receiver to not have to wait on each other
we'll see this in the next file, 03
this is known as creating a "buffered" channel
*/