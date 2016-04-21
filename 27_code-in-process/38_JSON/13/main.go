package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type StockData struct {
	Returns []float64 `json:"returns"`
}

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var obj StockData
	err = json.NewDecoder(f).Decode(&obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
