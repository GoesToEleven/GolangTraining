package main

import (
	"fmt"
	"time"
)

var workerId int = 0
var publisherId int = 0

func main() {
	input := make(chan string)
	go WorkerProcess(input)
	go WorkerProcess(input)
	go WorkerProcess(input)
	go Publisher(input)
	go Publisher(input)
	go Publisher(input)
	go Publisher(input)
	time.Sleep(1 * time.Millisecond)
}

// Publishers push data into a channel
func Publisher(out chan string) {
	publisherId += 1
	thisId := publisherId
	dataId := 0
	for {
		dataId += 1
		fmt.Printf("Publisher %d is pushing data\n", thisId)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisId, dataId)
		out <- data
	}
}

func WorkerProcess(in <-chan string) {
	workerId += 1
	thisId := workerId
	for {
		fmt.Printf("%d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisId, input)
	}
}

/*
CHALLENGE #1
Is this fan out?
My Answer:
Yes and no.
-- YES
Are we "fanning out" work? Yes. We've launched several goroutines that are simultaneously publishing a message onto our channel. The golang blog says, "Fan out means you have multiple functions reading from the same channel until that channel is closed." Here we do have multiple functions reading from the same channel. So, okay, we're fanning out.
-- NO
I should also note, however, that the "fan out" process can produce multiple channels. How are we ever going to "fan in" from multiple channels if we haven't produced those channels in the "fan out" process? Remember the "square output" example from the golang blog. The "fan out" process produced a channel each time a func read from the same input channel. Ultimately, as there is no way to "fan in" if we haven't "fanned out" onto multiple channels, I would say that, no, this code does not "fan out"


CHALLENGE #2
Is this fan in?
No.
What is being "fanned in" here? We have launched several goroutines of the same function: workerProcess. What do those goroutines do? They are all reading from an unbuffered channel. If there was a tremendous amount of processing that each "workerProcess" func executed, then all three of the "workerProcess" funcs could be processing in parallel: pulling values off the channel and processing them. There is no "fanning in" though here. Remember what the golang blog describes fan in: "A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed." We don't have many channels here converging into one channel.

*/

/*
code source:
https://gist.github.com/atedja/bba9ee75835702e1c8bc
*/
