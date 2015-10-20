package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/nu7hatch/gouuid"
	"io"
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
	cookie := createCookie()
	http.SetCookie(res, cookie)
	io.WriteString(res, `<!DOCTYPE html>
		<html>
		  <body>
			<p>`+fmt.Sprint(cookie.String()+cookie.Value)+`</p>
		  </body>
		</html>`)
}

func createCookie() *http.Cookie {
	hashKey := securecookie.GenerateRandomKey(32)
	s := securecookie.New(hashKey, blockKey)
	sessionID, _ := uuid.NewV4()
	log.Println("UUID: ", sessionID)
	encoded, _ := s.Encode("sessionID", sessionID)
	cookie := &http.Cookie{
		Name:  "sessionID",
		Value: encoded,
	}
	return cookie
}
