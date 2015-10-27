package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		val := req.URL.Query().Get(key)
		io.WriteString(res, "Do my search:"+val)
	})
	http.ListenAndServe(":9000", nil)
}

// visit this page:
// http://localhost:9000/?q="dog"
