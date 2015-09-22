// go get uuid
// https://github.com/nu7hatch/gouuid
// NewV4
package main

import (
	"fmt"
	"net/http"
	"github.com/nu7hatch/gouuid"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		// cookie is not set
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-id",
				Value: id.String(),
			}
			http.SetCookie(res, cookie)
		}
		fmt.Println(cookie)

	})
	http.ListenAndServe(":9000", nil)
}
