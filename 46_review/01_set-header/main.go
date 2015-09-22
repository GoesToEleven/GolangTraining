package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/plain")
		res.Header()["Content-Type"] = []string{"text/plain"} // same as line 10
		fmt.Fprint(res, "Dog")
	})

	http.ListenAndServe(":9000", nil)
}
