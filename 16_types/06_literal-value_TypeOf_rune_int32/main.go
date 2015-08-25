package main

import (
	"fmt"
	"reflect"
)

func main() {
	rune := 'd'
	fmt.Printf("%T\n", rune)
	fmt.Println(reflect.TypeOf(rune)," - ",rune)
}

// rune        alias for int32