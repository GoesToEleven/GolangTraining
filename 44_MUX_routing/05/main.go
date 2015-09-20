package main

import (
	"io"
	"net/http"
	"strings"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}
	io.WriteString(res, `
	Dog Name: <strong>`+dogName+`</strong><br>
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var catName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		catName = fs[2]
	}
	io.WriteString(res, `
	Cat Name: <strong>`+catName+`</strong><br>
	<img src="https://upload.wikimedia.org/wikipedia/commons/0/06/Kitten_in_Rizal_Park%2C_Manila.jpg">
	`)
}

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":9000", mux)
}
