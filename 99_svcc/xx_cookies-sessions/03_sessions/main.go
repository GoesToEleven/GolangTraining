package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		// cookie is not set
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-id",
				Value: id.String() + " email=jon@email.com" + " JSON data" + " Whatever",
			}
			http.SetCookie(res, cookie)
		}
		fmt.Println(cookie)

	})
	http.ListenAndServe(":9000", nil)
}

// go get uuid
// https://github.com/nu7hatch/gouuid
// NewV4
