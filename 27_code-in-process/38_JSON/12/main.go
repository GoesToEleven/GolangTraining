package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Anything interface{}

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var obj map[string]interface{}
	err = json.NewDecoder(f).Decode(&obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
