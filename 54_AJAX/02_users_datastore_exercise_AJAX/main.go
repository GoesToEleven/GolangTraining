package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Profile struct {
	Email     string
	FirstName string
	LastName  string
	Age       int
}

func init() {
	router := httprouter.New()
	router.GET("/", showIndex)
	router.GET("/profile", showProfile)
	router.GET("/api/profile", getAPIProfile)
	router.POST("/api/profile", updateAPIProfile)
	http.Handle("/", router)
}

func getAPIProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	var profile Profile
	err := datastore.Get(ctx, key, &profile)
	if err != nil {
		profile.Email = u.Email
	}
	json.NewEncoder(res).Encode(profile)
}

func updateAPIProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var profile Profile
	json.NewDecoder(req.Body).Decode(&profile)

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	_, err := datastore.Put(ctx, key, &profile)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func showIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "/profile", 302)
}

func showProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(res, "templates.gohtml", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
