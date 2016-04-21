package main

import (
	"net/http"
)

func setCookie(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "COOKIE MONSTER",
	})
}

func main() {
	http.HandleFunc("/", setCookie)
	http.ListenAndServe(":9000", nil)
}
