package main

import (
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/cloud/datastore"
	"encoding/json"
	"google.golang.org/appengine/log"
)

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/login", Login)
	r.GET("/signup", Signup)
	r.GET("/api/checkUserName", checkUserName)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
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
	var Possibility string
	err := json.NewDecoder(req.Body).Decode(&Possibility)
	if err != nil {
		log.Errorf(ctx, "error decoding name possibility", err)
		return
	}
	q, err := datastore.NewQuery("Users").Filter("Username =", Possibility).Count(ctx)
	if err != nil {
		log.Errorf(ctx, "error running query", err, " - ", Possibility)
		return
	}
	if q == 1 {
		err := json.NewEncoder(res).Encode("true")
		if err != nil {
			log.Errorf(ctx, "error encoding username possibility response - one match", err, Possibility)
		}
	} else if q > 1 {
		log.Errorf(ctx, "hmm, we had more than one username", Possibility)
	} else {
		err := json.NewEncoder(res).Encode("false")
		if err != nil {
			log.Errorf(ctx, "error encoding username possibility response - no matches", err, Possibility)
		}
	}
}