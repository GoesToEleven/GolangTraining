package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("no-file.txt"); err != nil {
		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
