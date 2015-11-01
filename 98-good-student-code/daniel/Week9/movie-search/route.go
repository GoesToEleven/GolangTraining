package search

import "net/http"

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/addMovie", handleAdd)
	http.HandleFunc("/search", handleSearch)
}
