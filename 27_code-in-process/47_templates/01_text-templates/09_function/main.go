package main

import (
	"log"
	"os"
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
		"mycustomfunc": func() string {
			return "This should work"
		},
	})
	tpl, err = tpl.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, Page{
		Title: "My Title",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
