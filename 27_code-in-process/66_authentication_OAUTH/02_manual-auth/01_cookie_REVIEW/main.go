package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("logged-in")

	// no cookie
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "logged-in",
			Value: "0",
		}
	}

	// check log in: password entered == "secret"?
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "1",
			}
		}
	}

	// if logout, then logout and destroy cookie
	if req.URL.Path == "/logout" {
		cookie = &http.Cookie{
			Name:   "logged-in",
			Value:  "0",
			MaxAge: -1,
		}
	}

	http.SetCookie(res, cookie)

	// create string with html for response
	var html string

	// not logged in
	if cookie.Value == strconv.Itoa(0) {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1>LOG IN</h1>
			<form method="post" action="/">
				<h3>User name</h3>
				<input type="text" name="userName">
				<h3>Password</h3>
				<input type="text" name="password">
				<br>
				<input type="submit">
				<input type="submit" name="logout" value="logout">
			</form>
			</body>
			</html>`
	}

	// logged in
	if cookie.Value == strconv.Itoa(1) {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="/logout">LOG OUT</a></h1>
			</body>
			</html>`
	}

	io.WriteString(res, html)
}
