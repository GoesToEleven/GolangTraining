package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func following(follower, followee string, req *http.Request) (bool, error) {
	ctx := appengine.NewContext(req)
	userKey := datastore.NewKey(ctx, "Users", follower, 0, nil)
	x, err := datastore.NewQuery("Follows").Ancestor(userKey).Filter("Following =", followee).Count(ctx)
	return x > 0, err
}
