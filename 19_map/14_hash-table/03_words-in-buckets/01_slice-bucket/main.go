package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice of slice of string to hold slices of words
	buckets := make([][]string, 12)

	// code here has been updated from the recording
	// see below for explanation

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	// Print the words in one of the buckets
	// fmt.Println(buckets[6])
	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// comment out the above, then uncomment the below
	// a more uneven distribution
	// return len(word) % buckets
}

/*
UPDATED CODE
Up above, the code has been updated from the recording
I used to have this ...
		buckets = append(buckets, []string{})
... and changed it to this ...
		buckets[i] = []string{}

REASON:

This line of code ...
	buckets := make([][]string, 12)
... creates a slice with len and cap equal to 12. I can now access each of the twelve positions in the slice by index and assign values to them. If I "append" to this slice, like this ....
		buckets = append(buckets, []string{})
... I am adding another twelve positions to my slice; my len increases to 24 and my cap increases to 24. This is unnecessary. I can, instead, just direclty begin accessing the first twelve positions in my slice ... and that's why I changed the code to this ...
		buckets[i] = []string{}

EVEN MORE EXPLANATION

You don't even need this entire chunk of code ...

	for i := 0; i < 12; i++ {
		buckets[i] = []string{}
	}

... as this code ...

	buckets := make([][]string, 12)

... creates a slice holding a []string, but it doesn't yet have a len or cap, so later I use append which is how you add an item to a slice in a position that does not yet have an item (beyond its current len).

Thank you to Lee Trent for pointing this out!
*/
