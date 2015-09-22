package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

func main() {
	tpls := template.New("templates")

	tpls, err := tpls.ParseFiles("main.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse templates", err, err.Error())
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("my-cookie")

		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "my-cookie",
				Value: "LOGGED OUT",
			}
		}

		userName := req.FormValue("userName")
		password := req.FormValue("password")
		if userName == "You" && password == "Me" {
			cookie.Value = "Logged In"
		}

		logout := req.FormValue("logout")
		fmt.Println(logout)
		if logout == "logout" {
			cookie.Value = "Logged Out"
		}

		http.SetCookie(res, cookie)

		// execute template
		err = tpls.ExecuteTemplate(res, "main.gohtml", cookie.Value)
		if err != nil {
			log.Fatalln("couldn't respond", err, err.Error())
		}
	})

	http.ListenAndServe(":9000", nil)
}
