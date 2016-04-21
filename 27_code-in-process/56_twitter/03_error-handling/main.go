package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
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
