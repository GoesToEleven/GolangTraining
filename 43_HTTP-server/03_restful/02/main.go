package main

import (
	"io"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, "<strong>kitty kitty kitty<strong>")
	case "/dog":
		io.WriteString(res, "<strong>doggy doggy doggy<strong>")
	}
}

func main() {

	var h myHandler
	http.ListenAndServe(":9000", h)
}
