package main

import (
	"html/template"
	"net/http"
	"strconv"

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
	router.POST("/profile", updateProfile)
	http.Handle("/", router)
}

func showIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "/profile", 302)
}

func showProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	tpl, err := template.ParseFiles("templates/templates.gohtml")
	if err != nil {
		panic(err)
	}

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	var profile Profile
	err = datastore.Get(ctx, key, &profile)
	if err != nil {
		profile.Email = u.Email
	}

	err = tpl.ExecuteTemplate(res, "edit-form", profile)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func updateProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	age, _ := strconv.Atoi(req.FormValue("age"))
	profile := Profile{
		Email:     u.Email,
		FirstName: req.FormValue("firstname"),
		LastName:  req.FormValue("lastname"),
		Age:       age,
	}
	_, err := datastore.Put(ctx, key, &profile)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
	http.Redirect(res, req, "/profile", 302)
}
