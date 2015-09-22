package main

import (
	"encoding/csv"
	"os"
	"net/http"
	"html/template"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		f, err := os.Open("../../../resources/table.csv")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer f.Close()

		rdr := csv.NewReader(f)
		rows, err := rdr.ReadAll()
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		records := make([]record, 0, len(rows))

		for i, row := range rows {
			if i == 0 {
				continue
			}
			record := makeRecord(row)
			records = append(records, record)
		}

		tpl, err := template.ParseFiles("csv.gohtml")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		err = tpl.Execute(res, records)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	})

	http.ListenAndServe(":9000", nil)
}