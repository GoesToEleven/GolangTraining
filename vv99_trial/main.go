package main

import (
	"html/template"
	"log"
	"net/http"
)

var err error
var tpl *template.Template

type user struct {
	name     string
	password string
}

var u1 user

func main() {
	tpl, err = tpl.ParseGlob("assets/templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", u1)
}

func login(res http.ResponseWriter, req *http.Request) {

	// PROCESS FORM SUBMISSION
	if req.Method == "POST" {
		password := req.FormValue("password")
		username := req.FormValue("userName")
		u1 = user{name: username, password: password}
		log.Println(u1)
		// redirect to main page
		http.Redirect(res, req, "/", 302)
	}

	// Execute template
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}