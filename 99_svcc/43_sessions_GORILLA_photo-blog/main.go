package main

import (
	"github.com/gorilla/sessions"
	"io"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func authenticate(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")

	// check log in
	if req.Method == "POST" &&
	req.FormValue("password") != "" {
		password := req.FormValue("password")
		if password == "secret" {
			session.Values["loggedin"] = "true"
			fmt.Println("LOGGED IN: ", session.Values["loggedin"])
		}
	}


	// add data
	if req.Method == "POST" &&
	req.FormValue("data") != "" {
		fmt.Println("UPLOAD")
		src, hdr, err := req.FormFile("data")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()
		// create a new file
		wd, _ := os.Getwd()
		imgDir := "assets/imgs/"
		path := filepath.Join(wd, imgDir, hdr.Filename)
		dst, err := os.Create(path)
		fmt.Println(dst)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()
		// copy the uploaded file into the new file
		io.Copy(dst, src)
		// add image to session
		var data []string
		jsonData := session.Values["data"]
		fmt.Printf("Type jsonData: %T\n", jsonData)
		if jsonData != nil {
			json.Unmarshal([]byte(jsonData.(string)), &data)
		}
		data = append(data, hdr.Filename)
		bs, _ := json.Marshal(data)
		session.Values["data"] = string(bs)
		fmt.Println("cookie data: ", session.Values["data"])
	}

	// if logout, then logout
	if req.URL.Path == "/logout" {
		session.Values["loggedin"] = "false"
	}

	var html string

	// not logged in
	if session.Values["loggedin"] == "false" ||
	session.Values["loggedin"] == nil {
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
			<h1>ADD PHOTO</h1>
			<form method="POST" action="http://localhost:9000/" enctype="multipart/form-data">
				<input type="file" name="data" id="data">
				<input type="submit">
			</form>
    		<p>`+
				fmt.Sprint("cookie data: ", session.Values["data"]) +
			`</p>
			</body>
			</html>`
	}

	session.Save(req, res)
	io.WriteString(res, html)
}

func main() {
	http.HandleFunc("/", authenticate)
	http.ListenAndServe(":9000", nil)
}
