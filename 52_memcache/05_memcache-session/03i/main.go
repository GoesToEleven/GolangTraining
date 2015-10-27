package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("sessionid")
	if cookie == nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: id.String(),
		}
		http.SetCookie(res, cookie)
	}

	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, cookie.Value)
	if item == nil {
		item = &memcache.Item{
			Key:   cookie.Value,
			Value: []byte("???"),
		}
	}

	fmt.Fprintln(res, item)
}
