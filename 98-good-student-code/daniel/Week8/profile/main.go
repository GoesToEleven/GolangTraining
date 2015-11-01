package profile

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/createProfile", createProfile)
}

// Profile holds user profiles
type Profile struct {
	Email     string
	FirstName string
	LastName  string
}

func handle(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	var profile Profile
	err := datastore.Get(ctx, key, &profile)
	if err == datastore.ErrNoSuchEntity {
		http.Redirect(res, req, "/createProfile", http.StatusSeeOther)
		return
	} else if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	tpl, err := template.ParseFiles("viewProfile.gohtml")
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	err = tpl.Execute(res, &profile)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}

func createProfile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	if req.Method == "POST" {
		u := user.Current(ctx)
		profile := Profile{
			Email:     u.Email,
			FirstName: req.FormValue("firstname"),
			LastName:  req.FormValue("lastname"),
		}
		key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
		_, err := datastore.Put(ctx, key, &profile)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	f, err := os.Open("createProfile.gohtml")
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	io.Copy(res, f)
}
