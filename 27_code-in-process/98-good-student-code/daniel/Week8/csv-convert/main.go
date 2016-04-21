package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func readHeader(row []string) []string {
	headers := []string{}
	for _, v := range row {
		headers = append(headers, v)
	}
	return headers
}

func readRow(row []string, header []string) (map[string]string, error) {
	if len(row) != len(header) {
		return nil, errors.New("row and header do not have the same lengths")
	}
	newRow := map[string]string{}
	for i := range row {
		newRow[header[i]] = row[i]
	}
	return newRow, nil
}

func readCsv(filename string) ([]map[string]string, error) {
	f, err := os.Open(os.Args[1])
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	dataBytes, err := rdr.ReadAll()
	if err != nil {
		return nil, err
	}

	data := []map[string]string{}

	headers := readHeader(dataBytes[0])

	for _, v := range dataBytes[1:] {
		newRow, err := readRow(v, headers)
		if err != nil {
			return nil, err
		}
		data = append(data, newRow)
	}

	return data, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage:", os.Args[0], "<csv filename>")
	}

	w, err := os.Create("table.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer w.Close()

	data, err := readCsv(os.Args[1])

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
