package main

import (
	"fmt"
	"github.com/goestoeleven/GolangTraining/05_scope/14_invalid-capitalization/characters"
)

func main() {
	fmt.Println(characters.Sherrif())
	fmt.Println(characters.SherrifName) // valid - can access the variable
	fmt.Println(characters.SideKick) // invalid - cannot access the variable
}