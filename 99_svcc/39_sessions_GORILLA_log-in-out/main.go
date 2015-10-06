package main

import (
	"github.com/gorilla/sessions"
	"io"
	"net/http"
	"fmt"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func authenticate(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")

	// check log in
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			session.Values["loggedin"] = "true"
		}
	}

	// if logout, then logout
	if req.URL.Path == "/logout" {
		session.Values["loggedin"] = "false"
	}

	session.Save(req, res)
	var html string

	// not logged in
	if session.Values["loggedin"] == "false" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1>LOG IN</h1>
			<form method="post" action="http://localhost:9000/">
				<h3>User name</h3>
				<input type="text" name="userName" id="userName">
				<h3>Password</h3>
				<input type="text" name="password" id="password">
				<br>
				<input type="submit">
				<input type="submit" name="logout" value="logout">
			</form>
			</body>
			</html>`
	} else {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="http://localhost:9000/logout">LOG OUT</a></h1>
			</body>
			</html>`
	}

	io.WriteString(res, html)
}

func main() {
	http.HandleFunc("/", authenticate)
	http.ListenAndServe(":9000", nil)
}
