package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("my-cookie")
		// there is no cookie
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "my-cookie",
				Value: "0",
			}
		}
		// there is a cookie
		count, _ := strconv.Atoi(cookie.Value)
		count++
		cookie.Value = strconv.Itoa(count)

		http.SetCookie(res, cookie)

		io.WriteString(res, cookie.Value)
	})
	http.ListenAndServe(":9000", nil)
}
