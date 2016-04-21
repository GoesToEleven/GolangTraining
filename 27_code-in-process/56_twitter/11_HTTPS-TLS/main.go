package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"io/ioutil"
	stdlog "log"
	"net/http"
)

type User struct {
	Email    string
	UserName string
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
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	stdlog.Println("REQUEST BODY: ", sbs)
	q, err := datastore.NewQuery("Users").Filter("UserName=", sbs).Count(ctx)
	stdlog.Println("ERR: ", err)
	stdlog.Println("QUANTITY: ", q)
	if err != nil {
		fmt.Fprint(res, "false")
		return
	}
	if q >= 1 {
		fmt.Fprint(res, "true")
	} else {
		fmt.Fprint(res, "false")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	NewUser := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: req.FormValue("password"),
	}
	key := datastore.NewIncompleteKey(ctx, "Users", nil)
	key, _ = datastore.Put(ctx, key, &NewUser)
	http.Redirect(res, req, "/", 302)
}

/*
FYI - good to read:

https://cloud.google.com/appengine/docs/go/config/appconfig

https://cloud.google.com/appengine/docs/using-custom-domains-and-ssl

*/
