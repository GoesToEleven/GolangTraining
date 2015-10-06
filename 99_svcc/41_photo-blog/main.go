package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"os"
	"io"
	"path/filepath"
	"strings"
)

var err error
var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Model struct {
	Pictures    []string
}

// []string of picture paths
var Data Model = getPicturePaths()

func main() {
	// Parse templates
	tpl, err = tpl.ParseGlob("assets/templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/logout", logout)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())


	// to generate cert and key:
	// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", Data)
}

func login(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")

	// PROCESS FORM SUBMISSION
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			session.Values["logged_in"] = true
		} else {
			http.Error(res, "invalid credentials", 401)
			return
			}
		// save session
		session.Save(req, res)
		// redirect to main page
		http.Redirect(res, req, "/", 302)
		return
		}
	// Execute template
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	delete(session.Values, "logged_in")
	session.Save(req, res)
	http.Redirect(res, req, "/", 302)
}

func upload(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")

	// redirect if not logged in
	if session.Values["logged_in"] == false ||
		session.Values["logged_in"] == nil {
		http.Redirect(res, req, "https://localhost:8080/login", 302)
		return
	}

	if req.Method == "POST" {
		// <input type="file" name="my-file">
		src, hdr, err := req.FormFile("my-file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()
		// create a new file
		path := "/Users/tm002/Documents/go/src/github.com/goestoeleven/GolangTraining/49_cookies-sessions/12_photo-blog/01i/assets/imgs/"
		dst, err := os.Create(path + hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()
		// copy the uploaded file into the new file
		io.Copy(dst, src)
		Data = getPicturePaths()
		http.Redirect(res, req, "/", 302)
		return
	}
	// Execute template
	tpl.ExecuteTemplate(res, "upload-file.gohtml", nil)
}

func getPicturePaths() Model {
	files := []string{}
	// OR: files := make([]string, 0)
	filepath.Walk("./", func(path string, fi os.FileInfo, err error) error {
		// skip directories
		if fi.IsDir() {
			return nil
		}
		// for windows: replace \ with /
		path = strings.Replace(path, "\\", "/", -1)
		// I only want .jpg files
		if strings.HasSuffix(path, ".jpg") {
			files = append(files, path)
		}
		return nil
	})
	return Model{Pictures: files}
}