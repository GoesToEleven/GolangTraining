package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	var err error

	tpl := template.New("tpl.gohtml")
	tpl = tpl.Funcs(template.FuncMap{
		"uppercase": func(str string) string {
			return strings.ToUpper(str)
		},
	})
	tpl, err := tpl.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, Page{
		Title: "My Title 2",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
