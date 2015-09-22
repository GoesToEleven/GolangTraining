package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"reflect"
)

func main() {
	tpls := template.New("templates")

	tpls, err := tpls.ParseFiles("login.gohtml", "logout.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse templates", err, err.Error())
	}

	http.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
		err := tpls.ExecuteTemplate(res, "login.gohtml", nil)
		if err != nil {
			log.Fatalln("couldn't respond", err, err.Error())
		}
	})

	http.HandleFunc("/logout/", func(res http.ResponseWriter, req *http.Request) {
		err := tpls.ExecuteTemplate(res, "logout.gohtml", nil)
		if err != nil {
			log.Fatalln("couldn't respond", err, err.Error())
		}
	})

	http.ListenAndServe(":9000", nil)
}
