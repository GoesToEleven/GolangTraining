package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type monument struct {
	id        string
	elevation float64
	location  string
}

type monumentInfo struct {
	columns   map[string]int
	monuments []monument
}

func loadMonuments(rdr io.Reader) (*monumentInfo, error) {
	info := new(monumentInfo)
	reader := csv.NewReader(rdr)
	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if i == 0 {
			info.loadHeader(record)
		} else {
			err = info.parseRow(record)
			if err != nil {
				return nil, err
			}
		}
	}
	return info, nil
}

func (mi *monumentInfo) loadHeader(row []string) {
	mi.columns = make(map[string]int, len(row))
	for i, v := range row {
		mi.columns[v] = i
	}
}

func (mi *monumentInfo) parseRow(row []string) error {
	info := mi.columns["Elevation in feet"]
	elevation, err := strconv.ParseFloat(row[info], 64)
	if err != nil {
		return err
	}
	location := row[mi.columns["Location"]]
	id := row[mi.columns["Station ID"]]
	mi.monuments = append(mi.monuments, monument{
		elevation: elevation,
		location:  location,
		id:        id,
	})
	return nil
}

func main() {
	file, err := os.Open("City_of_Champaign_GPS_Control_Points.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	monuments, err := loadMonuments(file)
	if err != nil {
		log.Fatalln(err)
	}

	var highest monument
	for _, v := range monuments.monuments {
		if v.elevation > highest.elevation {
			highest = v
		}
	}

	location := highest.location[1 : len(highest.location)-1]
	location = strings.Replace(location, " ", "", -1)

	fmt.Println(`<!DOCTYPE html>
<html>
  <body>
    <img src="https://maps.googleapis.com/maps/api/staticmap?zoom=16&maptype=satellite&size=700x700&markers=` + location + `">
    <h1>Station ` + string(highest.id) + `</h1>
    <h3>` + strconv.FormatFloat(highest.elevation, 'f', 2, 64) + `ft</h3>
  </body>
</html>`)
}
