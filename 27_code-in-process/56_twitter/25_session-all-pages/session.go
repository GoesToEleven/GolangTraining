package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func getSession(req *http.Request) *memcache.Item {
	cookie, err := req.Cookie("session")
	if err != nil {
		return &memcache.Item{}
	}

	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}
	}
	log.Infof(ctx, "%s", string(item.Value))
	return item
}
