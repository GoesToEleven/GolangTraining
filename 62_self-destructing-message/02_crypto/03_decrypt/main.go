package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
	"strings"
)

func main() {
	decrypted := "some message that you want to store / send securely"
	fmt.Println("BEFORE ENCRYPTION:", decrypted)

	var nonce [24]byte
	io.ReadAtLeast(rand.Reader, nonce[:], 24)
	var password [32]byte
	io.ReadAtLeast(rand.Reader, password[:], 32)
	encrypted := secretbox.Seal(nil, []byte(decrypted), &nonce, &password)
	// fmt.Printf("%T \n", encrypted)
	enHex := fmt.Sprintf("%x:%x", nonce[:], encrypted)
	fmt.Println("ENCRYPTED:", enHex)

	// decrypt
	var nonce2 [24]byte
	parts := strings.SplitN(enHex, ":", 2)
	if len(parts) < 2 {
		fmt.Errorf("expected nonce")
	}
	//get nonce
	bs, err := hex.DecodeString(parts[0])
	if err != nil || len(bs) != 24 {
		fmt.Errorf("invalid nonce")
	}
	copy(nonce2[:], bs)
	// get message
	bs, err = hex.DecodeString(parts[1])
	if err != nil {
		fmt.Errorf("invalid message")
	}
	// you need the password to open the sealed secret box
	msg, ok := secretbox.Open(nil, bs, &nonce2, &password)
	if !ok {
		fmt.Errorf("invalid message")
	}
	fmt.Println("AFTER DECRYPTING:", string(msg))
}
