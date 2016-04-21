package main

import (
	"io"
	"log"
	"os"
)

func main() {
	rdrs := make([]io.Reader, len(os.Args)-1)
	for i, v := range os.Args[1:] {
		f, err := os.Open(v)
		if err != nil {
			log.Fatalln("my program broke: ", err.Error())
		}
		defer f.Close()
		rdrs[i] = f
	}

	rdr := io.MultiReader(rdrs...)

	io.Copy(os.Stdout, rdr)
}
