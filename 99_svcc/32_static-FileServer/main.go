package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":9000", http.FileServer(http.Dir(".")))
}
