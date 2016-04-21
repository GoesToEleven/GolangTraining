package taskexample

import (
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <p>Nothing to do</p>
  </body>
</html>`)
}
