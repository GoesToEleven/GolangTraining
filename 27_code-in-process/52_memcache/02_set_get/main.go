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

	item1 := memcache.Item{
		Key:   "foo",
		Value: []byte("bar"),
	}

	memcache.Set(ctx, &item1)

	item, _ := memcache.Get(ctx, "foo")
	if item != nil {
		fmt.Fprintln(res, string(item.Value))
	}
}
