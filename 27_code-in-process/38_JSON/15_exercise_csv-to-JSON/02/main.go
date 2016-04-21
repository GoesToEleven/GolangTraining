package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type StockData struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

func toFloat(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

func main() {
	src, err := os.Open("../resources/table.csv")
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer src.Close()

	dst, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer dst.Close()

	rdr := csv.NewReader(src)

	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	listOfStockData := []StockData{}
	// put data into struct
	for _, row := range data {
		sd := StockData{
			Date:     row[0],
			Open:     toFloat(row[1]),
			High:     toFloat(row[2]),
			Low:      toFloat(row[3]),
			Close:    toFloat(row[4]),
			Volume:   toFloat(row[5]),
			AdjClose: toFloat(row[6]),
		}
		listOfStockData = append(listOfStockData, sd)
	}

	//convert to JSON
	err = json.NewEncoder(dst).Encode(listOfStockData)
	if err != nil {
		log.Fatalln("couldn't encode", err.Error())
	}
}
