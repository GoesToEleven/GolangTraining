package main

import (
	//	"API"
	"Memcache"
	"github.com/GoesToEleven/GolangTraining/56_twitter/18_abstract-API-Model/api"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/form/login", Login)
	r.GET("/form/signup", Signup)
	r.POST("/api/checkusername", API.CheckUserName)
	r.POST("/api/createuser", API.CreateUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	Memcache.Template(res, req, "Homepage", "home.html", tpl)
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	Memcache.Template(res, req, "Loginpage", "login.html", tpl)
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	Memcache.Template(res, req, "Signuppage", "signup.html", tpl)
}

/*
TO DO:
session
-memcache templates
- uuid in a cookie
--- https while logged in? - depends upon security required
- encrypt password on datastore?
--- never store an unencrypted password, so, resoundingly, YES
--- sha-256 fast hash value
- user memcache?
- datastore / memcache
session interface change
- change login button to logout when user logged in
post tweets
follow people
see tweets for everyone
see tweets for individual user
*/
