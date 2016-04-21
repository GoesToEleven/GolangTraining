package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

func main() {

	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
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
