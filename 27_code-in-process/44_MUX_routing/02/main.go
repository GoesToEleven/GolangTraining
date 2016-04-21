package main

import (
	"io"
	"net/http"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":9000", mux)
}

/*
if route ends in slash
/dog/
it includes anything beneath

if route ends in no-slash
/dog
it only includes that

*/
