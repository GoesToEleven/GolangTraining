package main

import (
	"fmt"
	"reflect"
)

func main() {
	phrase := 'invalid'
	fmt.Printf("%T\n", phrase)
	fmt.Println(reflect.TypeOf(phrase))
}