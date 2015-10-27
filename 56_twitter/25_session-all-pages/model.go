package main

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type templateData struct {
	User
	LoggedIn bool
}
