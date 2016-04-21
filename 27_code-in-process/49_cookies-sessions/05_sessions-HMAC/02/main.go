package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))

}

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		// cookie is not set
		if err != nil {
			//id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-id",
			}
		}

		if req.FormValue("email") != "" {
			cookie.Value = req.FormValue("email")
		}

		code := getCode(cookie.Value)
		cookie.Value = code + "|" + cookie.Value

		// this doesn't run
		// need more code added to work
		// just shown for example of how to do auth with HMAC

		http.SetCookie(res, cookie)

		io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <form method="POST">
    `+cookie.Value+`
      <input type="email" name="email">
      <input type="password" name="password">
      <input type="submit">
    </form>
  </body>
</html>`)

	})
	http.ListenAndServe(":9000", nil)
}
