package main

import "time"

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn      bool
	LoginFail     bool
	Tweets        []Tweet
	ViewingUser   string
	FollowingUser bool
	Following     []F
}

type Tweet struct {
	Msg      string
	Time     time.Time
	UserName string
}

type F struct {
	Following string
}
