package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"net/http"
	"time"
)

type Profile struct {
	Email    string
	Username string
}

type Tweet struct {
	Message  string
	Time     time.Time
	Username string
}

// get profile by username
func getProfileByUsername(req *http.Request, username string) (*Profile, error) {
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Profile").Filter("Username =", username).Limit(1)
	var profiles []Profile
	_, err := q.GetAll(ctx, &profiles)
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, fmt.Errorf("profile not found")
	}
	return &profiles[0], nil
}

// get profile by email
func getProfileByEmail(ctx context.Context, email string) (*Profile, error) {
	key := datastore.NewKey(ctx, "Profile", email, 0, nil)
	var profile Profile
	err := datastore.Get(ctx, key, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// get profile
func getProfile(ctx context.Context) (*Profile, error) {
	u := user.Current(ctx)
	return getProfileByEmail(ctx, u.Email)
}

// create profile
func createProfile(req *http.Request, profile *Profile) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Profile", profile.Email, 0, nil)
	_, err := datastore.Put(ctx, key, profile)
	return err
	// you can use memcache also to improve your consistency
}

// for eventual consistency thing
// assumption: user will show up in 10 seconds on datastore
func waitForProfile(req *http.Request, username string) error {
	deadline := time.Now().Add(time.Second * 10)
	for time.Now().Before(deadline) {
		_, err := getProfileByUsername(req, username)
		if err == nil {
			return nil
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

//// insert tweet
func putTweet(ctx context.Context, tweet *Tweet, email string) error {
	userKey := datastore.NewKey(ctx, "Profile", email, 0, nil)
	tweetKey := datastore.NewIncompleteKey(ctx, "Tweet", userKey)
	_, err := datastore.Put(ctx, tweetKey, tweet)
	if err != nil {
		return err
	}
	return nil
}

// get user tweets
func userTweets(ctx context.Context, email string) ([]Tweet, error) {
	var tweets []Tweet
	userKey := datastore.NewKey(ctx, "Profile", email, 0, nil)
	q := datastore.NewQuery("Tweet").Ancestor(userKey).Order("-Time").Limit(20)
	_, err := q.GetAll(ctx, &tweets)
	return tweets, err

}

// get recent tweets
func recentTweets(ctx context.Context) ([]Tweet, error) {
	var tweets []Tweet
	q := datastore.NewQuery("Tweet").Order("-Time").Limit(20)
	_, err := q.GetAll(ctx, &tweets)
	return tweets, err

}

//// delete tweet
//func delTweet(ctx context.Context, username string) (*Profile, error) {
//
//}
