package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	log.Flags()
	log.SetFlags(0)

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(os.Stdout, "Hello World")
	if err != nil {
		log.Fatalln(err)
	}
}
