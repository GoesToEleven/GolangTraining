package main

import (
	"io"
	"log"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		f, err := os.Open(v)
		if err != nil {
			log.Fatalln("my program broke: ", err.Error())
		}
		defer f.Close()

		io.Copy(os.Stdout, f)
	}
}
