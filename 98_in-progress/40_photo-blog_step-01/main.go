package main

import (
	"encoding/gob"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Model struct {
	Photos   []string
	loggedin bool
}

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	tpl, _ = template.ParseGlob("templates/*.gohtml")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/upload", upload)
	http.HandleFunc("/admin/logout", logout)
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs/"))))
	http.ListenAndServe(":9000", context.ClearHandler(http.DefaultServeMux))
}

func home(res http.ResponseWriter, req *http.Request) {
	var state Model
	session, _ := store.Get(req, "session")
	state.loggedin = session.Values["loggedin"] == "true"
	if err := json.Unmarshal(session.Values["photos"], &state.Photos); err != nil {
		panic(err)
	}
	state.Photos = session.Values["photos"]
	tpl.ExecuteTemplate(res, "index.gohtml", state)
}

func admin(res http.ResponseWriter, req *http.Request) {
	var state Model
	session, _ := store.Get(req, "session")
	password := req.FormValue("password")
	if password == "secret" {
		state.loggedin = true
		session.Values["loggedin"] = "true"
		session.Save(req, res)
		http.Redirect(res, req, "/admin/upload", 302)
		return
	}
	tpl.ExecuteTemplate(res, "admin_login.gohtml", nil)
}

func upload(res http.ResponseWriter, req *http.Request) {
	if state.loggedin != true {
		http.Redirect(res, req, "/admin", 302)
		return
	}
	// upload file
	if req.Method == "POST" {
		src, hdr, _ := req.FormFile("file")
		defer src.Close()

		fileName := hdr.Filename
		dst, _ := os.Create("imgs/" + fileName)
		defer dst.Close()

		io.Copy(dst, src)
		state.Photos = append(state.Photos, "imgs/"+fileName)
		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "admin_upload.gohtml", state)
}

func logout(res http.ResponseWriter, req *http.Request) {
	state.loggedin = false
	http.Redirect(res, req, "/", 302)
}
