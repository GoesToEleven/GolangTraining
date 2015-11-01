package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type information struct {
	date      string
	open      float64
	high      float64
	low       float64
	closeData float64
	volume    int64
	adjClose  float64
}

func parseHeader(row []string) map[string]int {
	columns := map[string]int{}
	for i, v := range row {
		columns[v] = i
	}
	return columns
}

func parseRow(columns map[string]int, row []string) (*information, error) {
	date := row[columns["Date"]]

	data := row[columns["Open"]]
	open, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	data = row[columns["High"]]
	high, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	data = row[columns["Low"]]
	low, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	data = row[columns["Close"]]
	closeData, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	data = row[columns["Volume"]]
	volume, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return nil, err
	}

	data = row[columns["Adj Close"]]
	adjClose, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	return &information{
		date:      date,
		open:      open,
		high:      high,
		low:       low,
		closeData: closeData,
		volume:    volume,
		adjClose:  adjClose,
	}, nil
}

func readTable(rdr io.Reader) ([]information, error) {
	csvReader := csv.NewReader(rdr)
	var columns map[string]int
	info := []information{}
	for i := 0; ; i++ {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if i == 0 {
			columns = parseHeader(row)
		} else {
			newData, err := parseRow(columns, row)
			if err != nil {
				return nil, err
			}
			info = append(info, *newData)
		}
	}
	return info, nil
}

func main() {
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	info, err := readTable(file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(`<!DOCTYPE html>
<html>
<head>
  <style>
    td {
      padding: 7px;
    }
  </style>
</head>
<body>
  <table border="1">
    <tr>
      <th>Date</th>
      <th>Open</th>
      <th>High</th>
      <th>Low</th>
      <th>Close</th>
      <th>Volume</th>
      <th>Adj Close</th>
    </tr>`)
	for _, v := range info {
		fmt.Printf(`    <tr>
      <td>%s</td>
      <td>%f</td>
      <td>%f</td>
      <td>%f</td>
      <td>%f</td>
      <td>%d</td>
      <td>%f</td>
    </tr>
`, v.date, v.open, v.high, v.low, v.closeData, v.volume, v.adjClose)
	}
	fmt.Println(`  </table>
</body>
</html>`)
}
