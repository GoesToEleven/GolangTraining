package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
	{
	"name": "Todd McLeod"
	}
	`

	var obj map[string]string

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	fmt.Println(obj["name"])
	fmt.Printf("%T\n", obj["name"])
}
