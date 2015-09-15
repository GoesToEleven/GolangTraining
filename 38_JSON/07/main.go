package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
	100
	`

	var obj interface{}

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	fmt.Printf("%T\n", obj)

	v, ok := obj.(float64)
	if !ok {
		v = 0
	}
	x := 100 + v

	fmt.Println(x)
}
