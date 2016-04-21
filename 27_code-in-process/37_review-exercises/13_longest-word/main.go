package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func longestWord(rdr io.Reader) string {
	longWord := ""
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > len(longWord) {
			longWord = word
		}
	}
	return longWord
}

func main() {
	srcFile, err := os.Open("moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	fmt.Println(longestWord(srcFile))

}

/*
get moby dick from terminal:
curl -O http://www.gutenberg.org/files/2701/old/moby10b.txt
*/
