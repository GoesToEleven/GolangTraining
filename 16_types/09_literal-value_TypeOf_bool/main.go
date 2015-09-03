package main

import (
	"fmt"
	"reflect"
)

func main() {
	content := true
	fmt.Printf("%T\n", content)
	fmt.Println(reflect.TypeOf(content))
}