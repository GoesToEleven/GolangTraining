package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func wordCount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		word = strings.Replace(word, ",", "", 1)
		word = strings.Replace(word, ".", "", 1)
		counts[word]++
	}
	return counts
}

func main() {
	srcFile, err := os.Open("moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	counts := wordCount(srcFile)
	fmt.Println(counts["whale"])

}

/*
get moby dick from terminal:
curl -O http://www.gutenberg.org/files/2701/old/moby10b.txt
*/
