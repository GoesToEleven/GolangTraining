package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("my-cookie")
		fmt.Println(cookie, err)

		http.SetCookie(res, &http.Cookie{
			Name:  "my-cookie",
			Value: "some value",
		})
	})
	http.ListenAndServe(":9000", nil)
}
