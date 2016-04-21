package main

import (
	"html/template"
	"net/http"
)

func getTemplate(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		panic(err)
	}
	tpl.Execute(res, req.URL.Path)
}

func main() {
	http.HandleFunc("/", getTemplate)
	http.ListenAndServe(":9000", nil)
}
