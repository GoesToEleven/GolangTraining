package main

import (
	"fmt"
)

func main() {
	myPhrase := "What is a good phrase?"
//	for key, value := range myPhrase {
//		fmt.Println(key, " - ", value, " - ", string(value))
//	}

	for i := 0; i < len(myPhrase); i++ {
		fmt.Printf("%d - %d - %q - %b\n", i, myPhrase[i], myPhrase[i], myPhrase[i])
	}
}