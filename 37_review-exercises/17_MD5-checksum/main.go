package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	fmt.Printf("%x\n", h.Sum(nil))
}
