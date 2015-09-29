package main

import (
	"io"
	"net/http"
)

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

		http.SetCookie(res, cookie)

		io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <form>
    `+cookie.Value+`
      <input type="email" name="email">
      <input type="submit">
    </form>
  </body>
</html>`)

	})
	http.ListenAndServe(":9000", nil)
}
