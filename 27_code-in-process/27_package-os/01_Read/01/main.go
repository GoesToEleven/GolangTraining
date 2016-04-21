package main

import (
	"os"
)

func main() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	bs := make([]byte, 5)
	src.Read(bs)
	dst.Write(bs)
}

// this is a limit reader
// we limit what is read
// see io.ReadFull for func similiar to (*File)Read
