package csvexample

import (
	"io"
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handleInput)
	http.HandleFunc("/madoff", handleOutput)
}

func handleInput(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <form method="GET" action="/madoff">
      <label>Symbol #1:
        <input type="text" name="symbol1">
      </label>
      <input type="submit">
    </form>
  </body>
</html>`)
}

func handleOutput(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	symbol1 := req.FormValue("symbol1")

	getData(ctx, symbol1)

}
