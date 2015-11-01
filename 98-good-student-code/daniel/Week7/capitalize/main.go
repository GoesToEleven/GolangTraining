package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage:", os.Args[0], "<file>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		text := scan.Text()
		first := text[0]
		rest := text[1:]
		result := strings.ToUpper(string(first)) + rest
		fmt.Println(result)
	}
}
