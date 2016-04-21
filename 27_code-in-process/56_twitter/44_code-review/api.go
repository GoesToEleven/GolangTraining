package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"io/ioutil"
	"net/http"
	"time"
)

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
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf(ctx, "error creating password: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	user := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: string(hashedPass),
	}
	key := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	key, err = datastore.Put(ctx, key, &user)
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}

	createSession(res, req, user)
	// redirect
	http.Redirect(res, req, "/", 302)
}

func loginProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", req.FormValue("userName"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.FormValue("password"))) != nil {
		// failure logging in
		var sd SessionData
		sd.LoginFail = true
		tpl.ExecuteTemplate(res, "login.html", sd)
		return
	} else {
		user.UserName = req.FormValue("userName")
		// success logging in
		createSession(res, req, user)
		// redirect
		http.Redirect(res, req, "/", 302)
	}
}

func createSession(res http.ResponseWriter, req *http.Request, user User) {
	ctx := appengine.NewContext(req)
	// SET COOKIE
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		Path:  "/",
		// twenty minute session:
		MaxAge: 60 * 20,
		//		UNCOMMENT WHEN DEPLOYED:
		//		Secure: true,
		//		HttpOnly: true,
	}
	http.SetCookie(res, cookie)

	// SET MEMCACHE session data (sd)
	json, err := json.Marshal(user)
	if err != nil {
		log.Errorf(ctx, "error marshalling during user creation: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
	}
	memcache.Set(ctx, &sd)
}

func logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("session")
	// cookie is not set
	if err != nil {
		http.Redirect(res, req, "/", 302)
		return
	}

	// clear memcache
	sd := memcache.Item{
		Key:        cookie.Value,
		Value:      []byte(""),
		Expiration: time.Duration(1 * time.Microsecond),
	}
	memcache.Set(ctx, &sd)

	// clear the cookie
	cookie.MaxAge = -1
	http.SetCookie(res, cookie)

	// redirect
	http.Redirect(res, req, "/", 302)
}

func tweetProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)
	if err != nil {
		log.Infof(ctx, "Attempt to post tweet from logged out user")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	}
	// declare a variable of type user
	// initialize user with values from memcache item
	var user User
	json.Unmarshal(memItem.Value, &user)
	// declare a variable of type tweet
	// initialize it with values
	log.Infof(ctx, user.UserName)
	tweet := Tweet{
		Msg:      req.FormValue("tweet"),
		Time:     time.Now(),
		UserName: user.UserName,
	}
	// put in datastore
	err = putTweet(req, &user, &tweet)
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	// redirect
	time.Sleep(time.Millisecond * 500) // This is not the best code, probably. Thoughts?
	http.Redirect(res, req, "/", 302)
}

func follow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	// get session
	memItem, err := getSession(req)
	if err != nil {
		log.Infof(ctx, "Attempt to follow from logged out user")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	}
	// declare a variable of type user
	// initialize user with values from memcache item
	var user User
	json.Unmarshal(memItem.Value, &user)
	// get the datastore key for the follower
	followerKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	// get a datastore key for the followee
	followeeKey := datastore.NewKey(ctx, "Follows", ps.ByName("user"), 0, followerKey)
	// the follower is following the followee
	// put this into the datastore
	_, err = datastore.Put(ctx, followeeKey, &F{ps.ByName("user"), user.UserName})
	if err != nil {
		log.Errorf(ctx, "error adding followee: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	// send the "you are being followed" email
	emailKey := datastore.NewKey(ctx, "Users", ps.ByName("user"), 0, nil)
	var u User
	err = datastore.Get(ctx, emailKey, &u)
	if err != nil {
		log.Errorf(ctx, "error getting followee's email user data: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	followedEmail(res, req, u.Email)
	// return to user account
	http.Redirect(res, req, "/user/"+ps.ByName("user"), 302)
}

func unfollow(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(req)
	// get session
	memItem, err := getSession(req)
	if err != nil {
		log.Infof(ctx, "Attempt to unfollow from logged out user")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	}
	// declare a variable of type user
	// initialize user with values from memcache item
	var user User
	json.Unmarshal(memItem.Value, &user)
	// get the datastore key for the follower
	followerKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	// get a datastore key for the followee
	followeeKey := datastore.NewKey(ctx, "Follows", ps.ByName("user"), 0, followerKey)
	// delete entry in datastore
	err = datastore.Delete(ctx, followeeKey)
	if err != nil {
		log.Errorf(ctx, "error deleting followee: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	http.Redirect(res, req, "/user/"+ps.ByName("user"), 302)
}
