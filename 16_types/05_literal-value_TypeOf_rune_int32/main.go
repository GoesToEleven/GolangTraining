package main

import (
	"fmt"
	"reflect"
)

func main() {
	rune := 'd'
	fmt.Printf("%T\n", rune)
	fmt.Println(reflect.TypeOf(rune))
}

// rune        alias for int32

/*
A rune literal represents a rune constant, an integer value identifying
a Unicode code point. A rune literal is expressed as one or more characters
enclosed in single quotes, as in 'x' or '\n'. Within the quotes, any character
may appear except newline and unescaped single quote. A single quoted character
represents the Unicode value of the character itself,
*/