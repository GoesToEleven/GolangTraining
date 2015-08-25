package main

import (
	"fmt"
	"reflect"
)

func main() {
	age := 30
	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(age))
}
