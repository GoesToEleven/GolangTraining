package main

import (
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "doggy doggy doggy")
	})

	mux.HandleFunc("/cat/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "catty catty catty")
	})

	http.ListenAndServe(":9000", mux)
}
