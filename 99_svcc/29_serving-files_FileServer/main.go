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
	if len(fs) >= 3 {
		dogName = fs[2]
	}
	// the image doesn't serve
	io.WriteString(res, `
	<h1>Dog Name: `+dogName+`</h1><br>
	<img src="/toby.jpg">
	`)
}

func main() {
	// FileServer returns a handler that serves HTTP requests
	// with the contents of the file system rooted at root.
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", upTown)
	http.ListenAndServe(":9000", nil)
}
