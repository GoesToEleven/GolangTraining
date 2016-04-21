package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(r io.Reader) (map[string](int), string, int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	words := map[string](int){}
	largestWord := ""
	numWords := 0
	for scanner.Scan() {
		wordsOrWord := strings.ToLower(scanner.Text())
		wordsOrWord = strings.Replace(wordsOrWord, "-", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "/", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "*", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, ".", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, ",", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "\"", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "'", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "!", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "?", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "_", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "=", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "+", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "&", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "%", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "^", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "(", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, ")", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "$", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "<", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, ">", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "{", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "}", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "[", " ", -1)
		wordsOrWord = strings.Replace(wordsOrWord, "]", " ", -1)
		for _, word := range strings.Fields(wordsOrWord) {
			size := len(word)
			words[word]++
			numWords++
			if size > len(largestWord) {
				largestWord = word
			}
		}
	}
	return words, largestWord, numWords
}

func main() {
	message, _ := os.Open("moby10b.txt")
	words, largestWord, numWords := count(message)
	fmt.Println("Number of the word \"the\":", words["the"])
	fmt.Println("Number of the word \"and\":", words["and"])
	fmt.Println("Number of the word \"whale\":", words["whale"])
	fmt.Println("Number of the word \"try\":", words["try"])
	fmt.Println("Number of the word \"succeed\":", words["succeed"])
	fmt.Printf("There are %d unique words in the document\n", len(words))
	fmt.Printf("There are %d total words in the document\n", numWords)
	fmt.Printf("The largest word is \"%s\", with a length of %d which appears %d times\n", largestWord, len(largestWord), words[largestWord])
}
