package main

import (
	"encoding/csv"
	"os"
	"net/http"
	"html/template"
)

func handle(path string, handleFunc func(res http.ResponseWriter, req *http.Request) error) {
	http.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		err := handleFunc(res, req)
		if err != nil {
			http.Error(res, err.Error(), 500)
		}
	})
}

func main() {
	handle("/", func(res http.ResponseWriter, req *http.Request) error {
		// read csv file
		f, err := os.Open("table.csv")
		if err != nil {
			return err
		}
		defer f.Close()

		// convert to slice of strings
		rdr := csv.NewReader(f)
		rows, err := rdr.ReadAll()
		if err != nil {
			return err
		}

		// convert to records
		records := make([]record, 0, len(rows))
		for i, row := range rows {
			if i == 0 {
				continue
			}
			record := makeRecord(row)
			records = append(records, record)
		}

		// render to template
		tpl, err := template.ParseFiles("csv.gohtml")
		if err != nil {
			return err
		}
		err = tpl.Execute(res, records)
		if err != nil {
			return err
		}

		return nil
	})

	http.ListenAndServe(":9000", nil)
}


