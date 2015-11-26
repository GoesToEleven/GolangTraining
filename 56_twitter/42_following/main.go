package main

import (
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", home)
	r.GET("/following", fing)
	r.GET("/user/:user", user)
	r.GET("/form/login", login)
	r.GET("/form/signup", signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	r.POST("/api/login", loginProcess)
	r.POST("/api/tweet", tweetProcess)
	r.GET("/api/logout", logout)
	r.GET("/api/follow/:user", follow)
	r.GET("/api/unfollow/:user", unfollow)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	tpl = template.New("roottemplate")
	tpl = tpl.Funcs(template.FuncMap{
		"humanize_time": humanize.Time,
	})

	tpl = template.Must(tpl.ParseGlob("templates/html/*.html"))
}

func home(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	//get tweets
	tweets, err := getTweets(req, nil)
	if err != nil {
		log.Errorf(ctx, "error getting tweets: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	// get session
	memItem, err := getSession(req)
	var sd SessionData
	if err == nil {
		// logged in
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	}
	sd.Tweets = tweets
	tpl.ExecuteTemplate(res, "home.html", &sd)
}

func user(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	user := User{UserName: ps.ByName("user")}
	//get tweets
	tweets, err := getTweets(req, &user)
	if err != nil {
		log.Errorf(ctx, "error getting tweets: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	// get session
	memItem, err := getSession(req)
	var sd SessionData
	if err == nil {
		// logged in
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		sd.ViewingUser = user.UserName
		sd.FollowingUser, err = following(sd.UserName, user.UserName, req)
		if err != nil {
			log.Errorf(ctx, "error running following query: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
	}
	sd.Tweets = tweets
	tpl.ExecuteTemplate(res, "user.html", &sd)
}

func fing(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	// get session
	memItem, err := getSession(req)
	if err != nil {
		log.Infof(ctx, "Attempt to see following from logged out user")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	}
	var sd SessionData
	if err == nil {
		// logged in
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	}
	// declare a variable of type user
	// initialize user with values from memcache item
	var user User
	json.Unmarshal(memItem.Value, &user)
	// Get followees
	// get the datastore key for the follower
	followerKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	var XF []F
	_, err = datastore.NewQuery("Follows").Ancestor(followerKey).Project("Following").GetAll(ctx, &XF)
	log.Errorf(ctx, "here is type %T \n and value %v", XF, XF)
	if err != nil {
		log.Errorf(ctx, "error getting followees: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	sd.Following = XF
	tpl.ExecuteTemplate(res, "follow.html", &sd)
}

func login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "login.html")
}

func signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "signup.html")
}
