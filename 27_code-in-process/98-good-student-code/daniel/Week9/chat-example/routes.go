package chat

import "net/http"

func init() {
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.Handle("/api/", newAPI("/api/"))
}
