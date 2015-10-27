package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type record struct {
	Date string
	Open float64
}

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return record{
		Date: row[0],
		Open: open,
	}
}

func main() {
	f, err := os.Open("../table.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		record := makeRecord(row)
		fmt.Println(record)

	}
}
