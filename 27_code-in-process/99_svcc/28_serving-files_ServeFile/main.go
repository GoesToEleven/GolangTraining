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
	<img src="/toby.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}

func main() {
	http.HandleFunc("/", upTown)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":9000", nil)
}
