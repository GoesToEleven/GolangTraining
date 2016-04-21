package main

import (
	"encoding/csv"
	//	"fmt"
	"encoding/json"
	"log"
	"os"
)

func main() {

	//open file
	src, err := os.Open("../../../resources/table.csv")
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer src.Close()

	// reader for csv file
	rdr := csv.NewReader(src)

	// read csv file
	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	// convert to JSON
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("couldn't marshall", err.Error())
	}

	// show
	os.Stdout.Write(b)
	//	fmt.Println(b)

}
