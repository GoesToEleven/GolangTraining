// BILL: THis is just a first pass.
package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	// BILL: Separate the stdlib from other imports.
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

// BILL
// Try to minimize package level variables. If you can't avoid them
// create a single struct type value.
// var s = struct {
//	tpl   *template.Template
//	store *sessions.Store
// }{
//	store: sessions.NewCookieStore([]byte("secret-password"))
// }

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	// BILL: You are ignoring an error. No!
	// BILL: If this fails call panic and find out why.
	tpl, _ = template.ParseGlob("assets/templates/*.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs", http.FileServer(http.Dir("./assets/imgs"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}

func index(res http.ResponseWriter, req *http.Request) {
	// BILL: Why are you ignoring an error. Which I am assuming
	// you are doing here.
	session, _ := store.Get(req, "session")
	
	// authenticate
	// BILL: Comments are full sentences with proper grammer. Same for all below.
	if session.Values["loggedin"] == "false" || session.Values["loggedin"] == nil {
		http.Redirect(res, req, "/login", 302)
		return
	}
	
	// upload photo
	src, hdr, err := req.FormFile("data")
	
	// BILL: NO. You must check the error first. If an error occurrs
	// you must ignore all other values.
	
	// BILL: Just realized you are not checking for the error.
	if req.Method == "POST" && err == nil {
		uploadPhoto(src, hdr, session)
	}
	
	// save session
	session.Save(req, res)
	
	// get photos
	// BILL: You don't need this variable. Call the
	// function inside the call to ExecuteTemplate.
	data := getPhotos(session)
	
	// execute template
	// I am sure this returns an error?
	tpl.ExecuteTemplate(res, "index.html", data)
}

func logout(res http.ResponseWriter, req *http.Request) {
	// BILL: Are you ignoring the error again?
	session, _ := store.Get(req, "session")
	
	session.Values["loggedin"] = "false"
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}

func login(res http.ResponseWriter, req *http.Request) {
	// BILL: Are you ignoring the error again?
	session, _ := store.Get(req, "session")
	if req.Method == "POST" && req.FormValue("password") == "secret" {
		session.Values["loggedin"] = "true"
		session.Save(req, res)
		http.Redirect(res, req, "/", 302)
		return
	}

	// execute template
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func uploadPhoto(src multipart.File, hdr *multipart.FileHeader, session *sessions.Session) {
	// BILL: I do not line defering the close here. Ths function does not
	// own the value. This is dangerous.
	defer src.Close()
	
	fName := getSha(src) + ".jpg"
	
	// BILL: Errors?
	wd, _ := os.Getwd()
	
	path := filepath.Join(wd, "assets", "imgs", fName)
	
	// BILL: Errors?
	dst, _ := os.Create(path)
	
	// BILL: Probably don't need a defer here since
	// this code will not panic and there is a single return.
	defer dst.Close()
	
	src.Seek(0, 0)
	io.Copy(dst, src)

	addPhoto(fName, session)
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func addPhoto(fName string, session *sessions.Session) {
	data := getPhotos(session)
	data = append(data, fName)
	
	// Error Checking?
	bs, _ := json.Marshal(data)
	session.Values["data"] = string(bs)
	
	// What if this fails, the user had no idea.
}

func getPhotos(session *sessions.Session) []string {
	// BILL: Yours
	// var data []string
	// jsonData := session.Values["data"]
	// if jsonData != nil {
	// 	json.Unmarshal([]byte(jsonData.(string)), &data)
	// }
	// return data
	
	// Use if statement for negative path testing when
	// you can.
	
	jsonData, exists := session.Values["data"]
	if !exists {
		return nil
	}
	
	var data []string
	json.Unmarshal([]byte(jsonData.(string)), &data)
	return data
}
