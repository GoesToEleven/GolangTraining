package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	src, err := os.Open("initial.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>>", line)
	}
}

/*
scanners allow us to interact with files line-by-line
*/
