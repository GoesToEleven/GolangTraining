package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	url, _ := user.LogoutURL(ctx, "/")
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)

}

func admin(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	url, _ := user.LogoutURL(ctx, "/")
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `Welcome ADMIN, %s! (<a href="%s">sign out</a>)`, u, url)

}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/admin/", admin)
}
