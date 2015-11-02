package main

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func profileHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "templates/profile.html")
}
func homeHandler(res http.ResponseWriter, req *http.Request) {
	log.Infof(appengine.NewContext(req), "path: %v", req.URL.Path)
	if req.URL.Path != "/" {
		profileHandler(res, req)
		return
	}
	http.ServeFile(res, req, "templates/home.html")
}
func loginHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(req)
	u := user.Current(c)

	profile, err := getProfile(c, u.Email)
	if err != nil || profile.Username == "" {

		if req.Method == "POST" {
			err = createProfile(c, &Profile{
				Username: req.FormValue("username"),
				Email:    u.Email,
			})
		} else {
			http.ServeFile(res, req, "templates/login.html")
			return
		}
	}
	http.SetCookie(res, &http.Cookie{Name: "logged_in", Value: "true"})
	http.Redirect(res, req, "/"+profile.Username, 302)
}

func logoutHandler(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{Name: "logged_in", Value: ""})
	http.Redirect(res, req, "/", 302)
}

func tweetHandler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	switch req.Method {
	case "GET":
		// get a list of tweetss
		screams, err := getScreams(ctx, req.FormValue("username"))
		if err != nil {
			log.Errorf(ctx, "error getting screams: %v", err)
			return
		}
		err = json.NewEncoder(res).Encode(screams)
		if err != nil {
			log.Errorf(ctx, "error marshalling todos: %v", err)
			return
		}

	case "POST":
		profile, err := getProfile(ctx, u.Email)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		// create a tweet
		var scream Scream
		err = json.NewDecoder(req.Body).Decode(&scream)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		scream.Time = time.Now()
		scream.Username = profile.Username
		scream.Email = u.Email
		err = createScream(ctx, &scream)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		json.NewEncoder(res).Encode(&scream)
		time.Sleep(time.Second * 3)

	case "DELETE":
		// delete a tweet
	default:
		http.Error(res, "method not allowed", 405)
	}
}
func followHandler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	switch req.Method {
	case "POST":
		var userName string
		err := json.NewDecoder(req.Body).Decode(&userName)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		err = follow(ctx, u.Email, userName)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		json.NewEncoder(res).Encode(true)
	case "GET":
		profile, err := getProfile(ctx, u.Email)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		json.NewEncoder(res).Encode(profile.Following)
	case "DELETE":
		// delete a follower
	default:
		http.Error(res, "method not allowed", 405)
	}
}
