package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("sessionid")
	fmt.Fprintln(res, cookie)
}
