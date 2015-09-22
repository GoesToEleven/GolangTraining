package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		http.ServeFile(res, req, "temp.txt")
	})

	http.ListenAndServe(":9000", nil)
}
