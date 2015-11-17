package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadAtLeast(rand.Reader, nonce[:], 24)
	fmt.Println(nonce)
}