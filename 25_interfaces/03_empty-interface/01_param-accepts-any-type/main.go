package main

import "fmt"

type Animal struct {
	sound string
}

type Dog struct {
	Animal
	friendly bool
}

type Cat struct {
	Animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := Dog{Animal{"woof"}, true}
	fifi := Cat{Animal{"meow"}, true}
	specs(fido)
	specs(fifi)
}
