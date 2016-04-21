package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
}
