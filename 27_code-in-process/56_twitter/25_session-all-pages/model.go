package main

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type sessionData struct {
	User
	LoggedIn bool
}
