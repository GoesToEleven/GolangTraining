package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

func main() {

	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
