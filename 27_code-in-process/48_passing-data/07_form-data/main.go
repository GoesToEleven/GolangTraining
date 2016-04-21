package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {

	// parse template
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// receive form submission
		fName := req.FormValue("first")
		lName := req.FormValue("last")
		fmt.Println("fName: ", fName)
		fmt.Println("[]byte(fName): ", []byte(fName))
		fmt.Println("typeOf: ", reflect.TypeOf(fName))

		// execute template
		err = tpl.Execute(res, Person{fName, lName})
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":9000", nil)
}
