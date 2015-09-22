package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("request made")
		cookie, err := req.Cookie("my-cookie")

		if err != nil {
			cookie = &http.Cookie{
				Name:  "my-cookie",
				Value: "1",
			}
			fmt.Println(cookie, err)
		} else {
			i, _ := strconv.Atoi(cookie.Value)
			newVal := i + 1
			cookie.Value = (strconv.Itoa(newVal))
			//			cookie.Value = (fmt.Sprint(newVal))
		}

		http.SetCookie(res, cookie)
		io.WriteString(res, `<html><body>`+cookie.Value+`</body></html>`)
	})

	http.ListenAndServe(":9000", nil)
}
