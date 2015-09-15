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

	fmt.Printf("%T", obj)
}
