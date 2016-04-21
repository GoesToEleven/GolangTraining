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

	x := 100 + obj.(float64)

	// conversion
	// if you have a concrete type, use conversion
	// eg, string(obj)
	// eg, float64(obj)

	// casting
	// if you have an interface, use the cast
	// obj.(float64)

	fmt.Println(x)
}
