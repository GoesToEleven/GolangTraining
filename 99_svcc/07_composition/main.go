package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "James Bond",
		Age:  23,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
