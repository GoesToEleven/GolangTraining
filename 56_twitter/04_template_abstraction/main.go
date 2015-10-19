package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template
var err error

func main() {
	tpl, err = template.ParseGlob("templates/html/*.html")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}