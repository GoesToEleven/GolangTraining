package main
import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
)

type state struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func parseState(record []string) (*state, error) {

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
		fmt.Printf(columns)
		break

	}

	// #3 do stuff for each row
	// #4 print out a table of information
}

/*

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type stateInformation struct {
	columns map[string]int
}

type state struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func (info *stateInformation) setColumns(record []string) {
	for idx, column := range record {
		info.columns[column] = idx
	}
}

func (info *stateInformation) parseState(record []string) (*state, error) {
	column := info.columns["id"]
	id, err := strconv.Atoi(record[column])
	if err != nil {
		return nil, err
	}
	name := record[info.columns["name"]]
	abbreviation := record[info.columns["abbreviation"]]
	censusRegionName := record[info.columns["census_region_name"]]
	return &state{
		id:               id,
		name:             name,
		abbreviation:     abbreviation,
		censusRegionName: censusRegionName,
	}, nil
}

func main() {
	// #1 open a file
	f, err := os.Open("state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	stateLookup := map[string]*state{}

	info := &stateInformation{}

	// #2 parse a csv file
	csvReader := csv.NewReader(f)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			info.setColumns(record)
		} else {
			state, err := info.parseState(record)
			if err != nil {
				log.Fatalln(err)
			}
			stateLookup[state.abbreviation] = state
		}
	}

	// state-information AL
	if len(os.Args) < 2 {
		log.Fatalln("expected state abbreviation")
	}
	abbreviation := os.Args[1]
	state, ok := stateLookup[abbreviation]
	if !ok {
		log.Fatalln("invalid state abbreviation")
	}

	fmt.Println(`
<html>
    <head></head>
    <body>
      <table>
        <tr>
          <th>Abbreviation</th>
          <th>Name</th>
        </tr>`)

	fmt.Println(`
        <tr>
          <td>` + state.abbreviation + `</td>
          <td>` + state.name + `</td>
        </tr>
    `)

	fmt.Println(`
      </table>
    </body>
</html>
    `)
}





*/