package main

import (
	"fmt"
)

func main() {
	// read NOTE ONE below
	c1 := incrementor("1")
	c2 := incrementor("2")
	// NOTE TWO
	// the flow of code continues here in main
	// incrementor is called
	// launches goroutines
	// returns channels which will eventually be closed
	// by the code that is running in the goroutines
	// launched by incrementor
	// program flow in main continues on down vertically here

	// NOTE THREE
	// you're passing into "fanIn" two channels
	// fanIn HAS TO pull values off of those channels
	// otherwise: DEADLOCK
	// fanIn will need to launch goroutines to pull those values
	c := fanIn(c1, c2)

	// NOTE SEVEN
	// we are held up here
	// pulling values off of c
	// until c is closed
	// and all values have been pulled from c
	// at which point ...
	for n := range c {
		fmt.Println(n)
	}

	// NOTE EIGHT
	// our program is done
	// flow of code exits out of main
	// program ends
}

func incrementor(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 20; i++ {
			// NOTE NINE
			// use Sprint here, or Sprintln, not Sprintf
			// this code is formatted for Sprintln
			// Sprintf would need this
			// c <- fmt.Sprintf("Process: %v, printing %v", s, i)
			c <- fmt.Sprint("Process: "+s+" printing:", i)
		}
		// NOTE ONE
		// incrementor
		// every time it's called
		// create a channel, put values on the channel
		// ******** important ********
		// have some other goroutine somewhere
		// pulling values off the channel
		// ***************************
		// close the channel
		// return that closed channel
		// these "incrementor" goroutines are off and running
		close(c)
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	done := make(chan bool)

	// NOTE FOUR
	// these goroutines will pull values off the "incrementor" channels
	// I'm defining the func to have a channel as a parameter and
	// I'm passing the channels in as an argument
	// as this is good practice to avoid different channels
	// accessing the same data and creating a race condition
	// not needed here, but good practice
	// and good to know about

	go func(x <-chan string) {
		for n := range x {
			c <- n
		}
		done <- true
	}(input1)

	go func(x <-chan string) {
		for n := range x {
			c <- n
		}
		done <- true
	}(input2)

	// NOTE FIVE
	// this will signal when we're done writing values to c
	go func() {
		<-done
		<-done
		close(c)
	}()

	// NOTE SIX
	// all of the above code
	// just flows straight through
	// goroutines are launched
	// and even though they're not done processing
	// program flow comes to here and this func returns
	// the channel c
	return c
}
