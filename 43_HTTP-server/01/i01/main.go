package main

import (
	"net/http"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}
