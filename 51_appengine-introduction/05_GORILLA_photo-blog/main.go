package blog

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	tpl, _ = template.ParseGlob("assets/templates/*.html")
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs", http.FileServer(http.Dir("./assets/imgs"))))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(res http.ResponseWriter, req *http.Request) {
	defer context.Clear(req)
	session, _ := store.Get(req, "session")
	// authenticate
	if session.Values["loggedin"] == "false" || session.Values["loggedin"] == nil {
		http.Redirect(res, req, "/login", 302)
		return
	}
	// upload photo
	src, hdr, err := req.FormFile("data")
	if req.Method == "POST" && err == nil {
		uploadPhoto(src, hdr, session)
	}
	// save session
	session.Save(req, res)
	// get photos
	data := getPhotos(session)
	// execute template
	tpl.ExecuteTemplate(res, "index.html", data)
}

func logout(res http.ResponseWriter, req *http.Request) {
	defer context.Clear(req)
	session, _ := store.Get(req, "session")
	session.Values["loggedin"] = "false"
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}

func login(res http.ResponseWriter, req *http.Request) {
	defer context.Clear(req)
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
	defer src.Close()
	fName := getSha(src) + ".jpg"
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "assets", "imgs", fName)
	dst, _ := os.Create(path)
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
	bs, _ := json.Marshal(data)
	session.Values["data"] = string(bs)
}

func getPhotos(session *sessions.Session) []string {
	var data []string
	jsonData := session.Values["data"]
	if jsonData != nil {
		json.Unmarshal([]byte(jsonData.(string)), &data)
	}
	return data
}
