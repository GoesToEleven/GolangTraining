package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Anything interface{}

func main() {
	jsonData := `
	"100"
	`

	var obj interface{}
	var f io.Reader
	json.NewDecoder(f).Decode(&obj)

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode([]int{1, 2, 3, 4})

	//bs, err := json.Marshal()
	fmt.Println(buf.String())

}

// this code doesn't run
// just captured in lecture
