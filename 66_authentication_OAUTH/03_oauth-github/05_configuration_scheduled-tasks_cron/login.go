package githubexample

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Go look at the log")
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "checking that log.infof prints to terminal")
}
