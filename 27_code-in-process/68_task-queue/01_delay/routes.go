package taskexample

import (
	"google.golang.org/appengine"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	delayedPuppy.Call(ctx)
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <p>Nothing to do</p>
  </body>
</html>`)
}
