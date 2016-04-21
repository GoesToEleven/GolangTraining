package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(res, "RequestURI: %v\n", req.RequestURI)
		fmt.Fprintf(res, "req.URL: %v\n", req.URL)
		fmt.Fprintf(res, "String(): %v\n", req.URL.String())
		fmt.Fprintf(res, "Path: %v\n", req.URL.Path)
		fmt.Fprintf(res, "RequestURI: %v\n", req.URL.RequestURI())
		fmt.Fprintf(res, "Opaque: %v\n", req.URL.Opaque)
		fmt.Fprintf(res, "RawPath: %v\n", req.URL.RawPath)
		fmt.Fprintf(res, "RawQuery: %v\n", req.URL.RawQuery)
		fmt.Fprintf(res, "User: %v\n", req.URL.User)
		fmt.Fprintf(res, "IsAbs(): %v\n", req.URL.IsAbs())
		fmt.Fprintf(res, "Scheme: %v\n", req.URL.Scheme)
	})

	http.ListenAndServe(":9000", nil)
}
