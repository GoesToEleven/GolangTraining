package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func renderTemplate(res http.ResponseWriter, name string, data interface{}) {
	err := tpl.ExecuteTemplate(res, name, data)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
