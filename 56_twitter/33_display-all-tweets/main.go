package main

import (
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
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	r.POST("/api/login", loginProcess)
	r.POST("/api/tweet", tweetProcess)
	r.GET("/api/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "home.html")
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "login.html")
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "signup.html")
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
