package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func putTweet(req *http.Request, user *User, tweet *Tweet) error {
	ctx := appengine.NewContext(req)
	userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	key := datastore.NewIncompleteKey(ctx, "Tweets", userKey)
	_, err := datastore.Put(ctx, key, tweet)
	return err
}

func getTweets(req *http.Request, user *User) ([]Tweet, error) {
	ctx := appengine.NewContext(req)

	var tweets []Tweet
	q := datastore.NewQuery("Tweets")

	if user != nil {
		// show tweets of a specific user
		userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
		q = q.Ancestor(userKey)
	}

	q = q.Order("-Time").Limit(20)
	_, err := q.GetAll(ctx, &tweets)
	return tweets, err
}
