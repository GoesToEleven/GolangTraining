package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")

	go http.ListenAndServe(":9999", nil)
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}

}

/*
for a self-signed certificate:

go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com

eg

go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
*/
