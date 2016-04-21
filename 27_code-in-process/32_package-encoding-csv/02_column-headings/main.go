package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// #1 open a file
	f, err := os.Open("../state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	csvReader := csv.NewReader(f)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		columns := make(map[string]int)
		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		}

		fmt.Println(columns)
		break
	}
	// #3 do stuff for each row
	// #4 print out a table of information
}
