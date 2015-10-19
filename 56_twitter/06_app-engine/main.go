/*
Our web app will be a micro-blogging site. It will only allow people to
share 140 characters of their thoughts per post. GL2U.

An example of a tweet could be:

GOLANG WEB APP TRAININGS from Silicon Valley Code @sv_code_camp 1 of 2:
https://youtu.be/qeREX9r20YQ 2 of 2: https://youtu.be/cIatklLmr5I
*/
package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template
var err error

// Init is the entry point for our web app. Init parses our templates.
// Init handles our routing and defines our end-points. As this is an app
// engine application, we are not allowed to use main. We therefore have
// changed our func main() to func init().
func init() {
	tpl, err = template.ParseGlob("templates/html/*.html")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	// We do not need ListenAndServe on app engine
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

// home handles everything coming into the root of our web app.
func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}