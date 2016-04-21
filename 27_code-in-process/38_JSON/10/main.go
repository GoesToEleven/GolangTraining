package main

import (
	"bytes"
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

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode([]int{1, 2, 3, 4})
	fmt.Println(buf)
}
