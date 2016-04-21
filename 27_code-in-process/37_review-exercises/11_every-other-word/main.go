package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	src, err := os.Open("initial.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	i := 0
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 0 && i%2 == 0 {
			fmt.Print(strings.ToUpper(word), " ")
		} else {
			fmt.Print(word, " ")
		}
		i++
	}
	fmt.Println()
}
