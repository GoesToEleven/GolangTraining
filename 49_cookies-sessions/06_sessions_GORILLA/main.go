package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"io"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, "session")
		if req.FormValue("email") != "" {
			// check password
			session.Values["email"] = req.FormValue("email")
		}
		session.Save(req, res)

		io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <form method="POST">
    `+fmt.Sprint(session.Values["email"])+`
      <input type="email" name="email">
      <input type="password" name="password">
      <input type="submit">
    </form>
  </body>
</html>`)

	})
	http.ListenAndServe(":9000", nil)
}
