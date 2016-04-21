package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func getSession(req *http.Request) (*memcache.Item, error) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return &memcache.Item{}, err
	}

	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		return &memcache.Item{}, err
	}
	return item, nil
}
