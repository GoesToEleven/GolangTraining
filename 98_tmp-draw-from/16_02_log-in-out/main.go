package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tpl := template.New("loginTemplate")
	tpl, err := tpl.ParseFiles("login.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse", err, err.Error())
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err := tpl.ExecuteTemplate(res, "login.gohtml", nil)
		if err != nil {
			log.Fatalln("couldn't respond", err, err.Error())
		}
	})

	http.ListenAndServe(":9000", nil)
}
