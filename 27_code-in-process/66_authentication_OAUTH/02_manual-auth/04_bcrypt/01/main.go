package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "mywifesnameandbirthday"
	bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	fmt.Println(bs)
	//	fmt.Println(string(bs))
	fmt.Printf("%X \n", bs)

	err := bcrypt.CompareHashAndPassword(bs, []byte(p))
	if err != nil {
		fmt.Println("Doesn't match")
	} else {
		fmt.Println("match")
	}
}
