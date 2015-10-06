package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func setCookie(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	// if cookie is not set
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + " email=jon@email.com" + " JSON data" + " Whatever",
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}

func main() {
	http.HandleFunc("/", setCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}

// go get uuid
// https://github.com/nu7hatch/gouuid
// NewV4
