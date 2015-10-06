package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Model struct {
	Photos []string
	loggedin bool
}

var state Model
var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.gohtml")

	filepath.Walk("imgs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path,"jpg"){
			state.Photos = append(state.Photos, path)
		}
		return nil
	})
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/upload", upload)
	http.HandleFunc("/admin/logout", logout)
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs/"))))
	http.ListenAndServe(":9000", nil)
}

func home(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", state)
}

func admin(res http.ResponseWriter, req *http.Request) {
	password := req.FormValue("password")
	if password == "secret" {
		state.loggedin = true
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
		state.Photos = append(state.Photos, "imgs/" + fileName)
		http.Redirect(res, req, "/", 302)
		return
	}
	tpl.ExecuteTemplate(res, "admin_upload.gohtml", state)
}

func logout(res http.ResponseWriter, req *http.Request) {
	state.loggedin = false
	http.Redirect(res, req, "/", 302)
}