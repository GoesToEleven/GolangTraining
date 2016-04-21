package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	item := loggedIn(req)
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "%v", item.Value)
	log.Infof(ctx, "%v", len(item.Value))
	if len(item.Value) > 0 {
		var td templateData
		json.Unmarshal(item.Value, &td)
		log.Infof(ctx, "%v", td)
		log.Infof(ctx, "%v", td.Email)
		log.Infof(ctx, "%v", td.LoggedIn)
		td.LoggedIn = true
		log.Infof(ctx, "%v", td.LoggedIn)
		tpl.ExecuteTemplate(res, "home.html", td)
	} else {
		memTemplate(res, req, "Homepage", "home.html")
	}
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	memTemplate(res, req, "Loginpage", "login.html")
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	memTemplate(res, req, "Signuppage", "signup.html")
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
