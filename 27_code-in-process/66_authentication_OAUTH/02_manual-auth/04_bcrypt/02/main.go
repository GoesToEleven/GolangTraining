package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "todd"
	bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	fmt.Printf("PASSWORD ONE: %x \n", bs)

	p2 := "todd"
	bs2, _ := bcrypt.GenerateFromPassword([]byte(p2), bcrypt.MinCost)
	fmt.Printf("PASSWORD TWO: %x \n", bs2)
}
