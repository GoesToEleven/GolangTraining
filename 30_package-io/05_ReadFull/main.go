package main

import (
	"io"
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
	io.ReadFull(src, bs)
	dst.Write(bs)

}

// this is a limit reader
// we limit what is read
// see (*File)Read (os package) for func similiar to io.ReadFull
