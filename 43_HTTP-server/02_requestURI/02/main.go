package main

import (
	"io"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, req.RequestURI)
}

func main() {

	var h myHandler
	http.ListenAndServe(":9000", h)
}
