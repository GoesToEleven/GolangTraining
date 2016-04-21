package main

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
)

func main() {
	decrypted := "some message that has not yet been encrypted."
	// the nonce must be unique for every message encrypted
	var nonce [24]byte
	io.ReadAtLeast(rand.Reader, nonce[:], 24)
	// the password must be unique for every message encrypted
	var password [32]byte
	io.ReadAtLeast(rand.Reader, password[:], 32)
	encrypted := secretbox.Seal(nil, []byte(decrypted), &nonce, &password)
	fmt.Println("-----DECRYPTED-----")
	fmt.Println(decrypted)
	fmt.Println("-----ENCRYPTED-----")
	fmt.Println("encrypted", encrypted)
	fmt.Println("len", len(encrypted))
	fmt.Println("string", string(encrypted))
	fmt.Println("nonce", nonce)
	fmt.Println("nonce", nonce[:])
	fmt.Printf("%x:%x \n", nonce[:], encrypted)
}
