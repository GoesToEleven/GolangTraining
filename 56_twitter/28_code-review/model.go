package main

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
}
