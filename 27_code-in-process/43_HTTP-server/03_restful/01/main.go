package main

import (
	"io"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	}
}

func main() {

	var h myHandler
	http.ListenAndServe(":9000", h)
}
