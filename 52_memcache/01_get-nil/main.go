package main

import (
	"fmt"
	"net/http"
)
import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, "some-key")
	fmt.Fprintln(res, item)
}
