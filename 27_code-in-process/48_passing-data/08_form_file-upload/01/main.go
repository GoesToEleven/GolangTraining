package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	// parse template
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// receive form submission
		if req.Method == "POST" {
			src, _, err := req.FormFile("file")
			if err != nil {
				panic(err)
			}
			defer src.Close()

			dst, err := os.Create(filepath.Join("./", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()

			io.Copy(dst, src)
		}

		// execute template
		err = tpl.Execute(res, nil)
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":9000", nil)
}
