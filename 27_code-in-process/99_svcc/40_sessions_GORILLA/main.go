package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"io"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func setSession(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if req.FormValue("email") != "" {
		session.Values["email"] = req.FormValue("email")
	}
	session.Save(req, res)

	io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <form method="POST">
    <p>`+fmt.Sprint(session.Values["email"])+`</p>
    <p>Enter Your Email:</p>
      <input type="email" name="email">
      <input type="submit">
    </form>
  </body>
</html>`)
}

func main() {
	http.HandleFunc("/", setSession)
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
