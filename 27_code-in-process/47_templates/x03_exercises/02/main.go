package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// function
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// execute template
		err = tpl.Execute(res, req.RequestURI)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// create server
	http.ListenAndServe(":9000", nil)
}
