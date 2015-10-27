package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"html/template"
	"io/ioutil"
	"net/http"
	//	"io"
	//	"bytes"
	//	"google.golang.org/appengine/memcache"
)

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string
}

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/form/login", Login)
	r.GET("/form/signup", Signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // maybe not needed b/c of schmidt router
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	log.Infof(ctx, "REQUEST BODY: %v", sbs)
	var user User
	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
	err = datastore.Get(ctx, key, &user)
	// if there is an err, there is NO user
	log.Infof(ctx, "ERR: %v", err)
	if err != nil {
		// there is an err, there is a NO user
		fmt.Fprint(res, "false")
		return
	} else {
		fmt.Fprint(res, "true")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	NewUser := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: req.FormValue("password"),
	}
	key := datastore.NewKey(ctx, "Users", NewUser.UserName, 0, nil)
	key, err := datastore.Put(ctx, key, &NewUser)
	// this is the only error checking I added; any others on this page needed?
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	http.Redirect(res, req, "/", 302)
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
