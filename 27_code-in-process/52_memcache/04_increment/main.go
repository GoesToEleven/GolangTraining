package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/user"
)

func index(res http.ResponseWriter, req *http.Request) {
	// gets rid of favicon.ico requests and any other requests
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	globalCount, _ := memcache.Increment(ctx, "GLOBAL", 1, 0)
	userCount, _ := memcache.Increment(ctx, u.Email+".COUNTER", 1, 0)

	fmt.Fprintln(res, "Global", globalCount)
	fmt.Fprintln(res, "User", userCount)

}

func init() {
	http.HandleFunc("/", index)
}
