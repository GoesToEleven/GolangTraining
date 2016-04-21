package main

import (
	"io"
	"net/http"
	"strings"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 2 {
		dogName = fs[1]
	}
	// the image doesn't serve
	io.WriteString(res, `
	<h1>Dog Name: `+dogName+`</h1><br>
	<img src="/resources/toby.jpg">
	`)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", upTown)
	http.ListenAndServe(":9000", nil)
}
