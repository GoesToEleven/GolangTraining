package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type dataType struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func main() {
	data := dataType{"Daniel", 22, "Student"}
	json.NewEncoder(os.Stdout).Encode(data)

	f, err := os.Open("data.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()
	var readData []float64
	err = json.NewDecoder(f).Decode(&readData)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(readData)
}
