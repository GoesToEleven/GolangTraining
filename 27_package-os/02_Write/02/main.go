package main

import (
	"log"
	"os"
)

func main() {

	dst, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln("error creating destination file: ", err.Error())
	}
	defer dst.Close()

	bs := []byte("Put some phrase here.")

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}
