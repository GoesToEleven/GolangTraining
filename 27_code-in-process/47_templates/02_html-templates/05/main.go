package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	var err error
	var tpl *template.Template
	tpl, err = tpl.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, Page{
		Title: "My Title 2",
		Body:  "hello world",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\n***************")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Page{
		Title: "My Title 2",
		Body:  "hello world",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\n***************")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", Page{
		Title: "My Title 2",
		Body:  "hello world",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
