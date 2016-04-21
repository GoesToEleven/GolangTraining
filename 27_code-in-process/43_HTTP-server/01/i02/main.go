package main

import (
	"io"
	"net/http"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World")
}

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}
