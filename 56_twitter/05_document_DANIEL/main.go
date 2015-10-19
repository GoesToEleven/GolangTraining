/*
Our web app will be a micro-blogging site. It will only allow people to
share 140 characters of their thoughts per post. GL2U.

An example of a tweet could be:

GOLANG WEB APP TRAININGS from Silicon Valley Code @sv_code_camp 1 of 2:
https://youtu.be/qeREX9r20YQ 2 of 2: https://youtu.be/cIatklLmr5I

Learn more about documenting your code:
https://golang.org/doc/effective_go.html#commentary
http://blog.golang.org/godoc-documenting-go-code

Use the godoc command to see your documentation:
https://godoc.org/golang.org/x/tools/cmd/godoc

Try these godoc commands:
godoc .
godoc -http=:6060
*/
package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template
var err error

// Main is the entry point for our web app. Main parses our templates.
// Main handles our routing and defines our end-points. Doc comments work
// best as complete sentences, which allow a wide variety of automated
// presentations. The first sentence should be a one-sentence summary
// that starts with the name being declared.
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

// home handles everything coming into the root of our web app.
func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}