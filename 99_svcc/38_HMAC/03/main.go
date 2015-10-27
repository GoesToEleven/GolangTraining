package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	// cookie is not set
	if err != nil {
		id, _ := uuid.NewV4()
		code := getCode(id.String())
		val := code + "|" + id.String()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: val,
		}
	}

	values := strings.Split(cookie.Value, "|")
	code := getCode(values[1])
	cookieCode := values[0]
	fmt.Fprintln(res, code)
	fmt.Fprintln(res, cookieCode)

	if code != cookieCode {
		fmt.Fprintln(res, "Cookie monsters says: someone's had their hands in my cookies!")
		cookie = &http.Cookie{
			Name:   "session-id",
			Value:  "0",
			MaxAge: -1,
		}
	}

	http.SetCookie(res, cookie)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
