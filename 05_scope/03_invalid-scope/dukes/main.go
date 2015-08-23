package main

import (
	"fmt"
	"github.com/goestoeleven/GolangTraining/05_scope/03_invalid-scope/characters"
)

func main() {
	fmt.Println(characters.Sherrif())
	fmt.Println(characters.sherrif) // not valid - can't access the variable
}