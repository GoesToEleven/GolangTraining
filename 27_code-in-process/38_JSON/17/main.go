package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
	// High, Low, Close
}

func main() {
	src, err := os.Open("../../../resources/table.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("table.json")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	rows, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	records := make([]Record, 0, len(rows))
	for _, row := range rows {
		date, _ := time.Parse("2006-01-01_this-does-not-compile", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	err = json.NewEncoder(dst).Encode(records)
	if err != nil {
		panic(err)
	}

}
