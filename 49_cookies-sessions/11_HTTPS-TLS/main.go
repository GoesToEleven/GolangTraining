package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func redir(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://127.0.0.1:10443/"+req.RequestURI, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")

	go http.ListenAndServe(":9999", http.HandlerFunc(redir))
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}

}
