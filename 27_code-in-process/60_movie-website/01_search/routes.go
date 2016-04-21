package movieinfo

import "net/http"

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/new-movie", handleNewMovie)
}
