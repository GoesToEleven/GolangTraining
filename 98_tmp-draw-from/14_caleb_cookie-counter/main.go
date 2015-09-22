package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("counter")
		// no cookie
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "counter",
				Value: "0",
			}
		}

		count, _ := strconv.Atoi(cookie.Value)
		count++
		cookie.Value = strconv.Itoa(count)

		http.SetCookie(res, cookie)

		io.WriteString(res, cookie.Value)

	})
	http.ListenAndServe(":9000", nil)
}


