package main

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/sessions"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))

	router := httprouter.New()
	router.GET("/", Home)
	router.GET("/signup", Signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.Handle("/", router)
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	session, _ := store.Get(req, "session")
}