package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		http.SetCookie(res, &http.Cookie{
			Name:  "my-cookie",
			Value: "some value",
		})
	})
	http.ListenAndServe(":9000", nil)
}
