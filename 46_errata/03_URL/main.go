package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(res, req.URL.Path)
	})

	http.ListenAndServe(":9000", nil)
}
