package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("sessionid")
	if cookie == nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: id.String(),
		}
		http.SetCookie(res, cookie)
	}

	fmt.Fprintln(res, cookie)
}
