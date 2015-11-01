package main

import (
	"net/http"
	"text/template"
)

func renderHome(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.ExecuteTemplate(res, "screamer", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func renderProfile(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/profile.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.ExecuteTemplate(res, "profile", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
