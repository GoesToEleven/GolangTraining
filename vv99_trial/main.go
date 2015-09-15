package main

import (
	"fmt"
)

func main() {
	var character rune = 'a'
	fmt.Println(character)
	fmt.Println(string(character))
	fmt.Printf("%v\n", character)
	fmt.Printf("%d\n", character)
	fmt.Printf("%#x\n", character)
	fmt.Printf("%T\n", character)

}
