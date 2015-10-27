package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type state struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func parseState(columns map[string]int, record []string) (*state, error) {
	id, err := strconv.Atoi(record[columns["id"]])
	name := record[columns["name"]]
	abbreviation := record[columns["abbreviation"]]
	censusRegionName := record[columns["census_region_name"]]
	if err != nil {
		return nil, err
	}
	return &state{
		id:               id,
		name:             name,
		abbreviation:     abbreviation,
		censusRegionName: censusRegionName,
	}, nil
}

func main() {
	// #1 open a file
	f, err := os.Open("../state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	csvReader := csv.NewReader(f)
	columns := make(map[string]int)

	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		} else {
			// #3 do stuff for each row
			state, err := parseState(columns, record)
			if err != nil {
				log.Fatalln(err)
			}
			// #4 print out each row
			log.Println(state)
		}
	}
}
