package main

import (
	"fmt"
	"reflect"
)

func main() {
	fname := "Jane"
	lname := `Doe`
	fmt.Printf("%T\n", fname)
	fmt.Println(reflect.TypeOf(lname))
}
