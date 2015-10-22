package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	// The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	key := "opensesame123456" // 16 bytes!
	block, _ := aes.NewCipher([]byte(key))
	fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), block.BlockSize)
	str := []byte("Hello World, and everyone else in the universe!")
	// 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16
	// encrypt
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)
	fmt.Printf("%s encrypted to %v\n", str, encrypted)
	// decrypt
	decrypter := cipher.NewCFBDecrypter(block, iv) // simple!
	decrypted := make([]byte, len(str))
	decrypter.XORKeyStream(decrypted, encrypted)
	fmt.Printf("%v decrypt to %s\n", encrypted, decrypted)
}

// https://www.socketloop.com/tutorials/golang-how-to-encrypt-with-aes-crypto
