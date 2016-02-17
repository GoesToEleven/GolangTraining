package main

import (
	"encoding/csv"
	"fmt"
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

	stateLookup := map[string]*state{}

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
			// #4 add each row to stateLookup map
			stateLookup[state.abbreviation] = state
		}
	}

	// #5 lookup the state
	if len(os.Args) < 2 {
		log.Fatalln("expected state abbreviation")
	}
	abbreviation := os.Args[1]
	state, ok := stateLookup[abbreviation]
	if !ok {
		log.Fatalln("invalid state abbreviation")
	}
	fmt.Println(state)
}

/*
at terminal:
go install

at terminal:
programName <state abbreviation>

for example:
step05_state-lookup CA

will show this:
&{5 California CA West}

*/
