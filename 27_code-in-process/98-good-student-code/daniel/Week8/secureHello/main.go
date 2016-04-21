package main

import (
	"io"
	"net/http"
)

func getPage(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World!")
}

func main() {
	http.HandleFunc("/", getPage)
	http.ListenAndServeTLS(":9000", "cert.pem", "key.pem", nil)
}
