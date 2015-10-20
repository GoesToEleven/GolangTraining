package main

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"log"
)

var blockKey []byte

func init() {
	blockKey = securecookie.GenerateRandomKey(32)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	hashKey := securecookie.GenerateRandomKey(32)
	s := securecookie.New(hashKey, blockKey)

	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("cookie-name", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(res, cookie)
	} else {
		log.Println("Nothing. ", err)
	}
}