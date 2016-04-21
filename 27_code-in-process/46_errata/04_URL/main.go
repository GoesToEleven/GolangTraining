package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(res, req.URL.EscapedPath())
		fmt.Fprintln(res, req.URL.RawQuery)
	})

	http.ListenAndServe(":9000", nil)
}
