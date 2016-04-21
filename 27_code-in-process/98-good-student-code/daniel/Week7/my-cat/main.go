package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough command line arguments")
	}

	readers := make([]io.Reader, len(os.Args)-1)
	for i, v := range os.Args[1:] {
		readFile, err := os.Open(v)
		if err != nil {
			log.Fatalln(err)
		}
		defer readFile.Close()
		readers[i] = readFile
	}
	rdr := io.MultiReader(readers...)

	io.Copy(os.Stdout, rdr)
}
