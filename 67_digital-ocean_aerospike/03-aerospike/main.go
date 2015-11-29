package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
)

type user struct {
	Username string
	Password string
	ID       string
}

var tpl *template.Template
var client *as.Client

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	f, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	client, err = as.NewClient("127.0.0.1", 3000)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/login", loginPage)
	router.POST("/login", login)
	router.GET("/logout", logout)
	router.GET("/create", createPage)
	router.POST("/create", create)
	err = http.ListenAndServe(":9000", router)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var username string
	c, err := req.Cookie("login")
	if err == nil {
		id := c.Value
		key, err := as.NewKey("bar", "users", id)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Printf("error creating key: %s\n", err.Error())
			return
		}
		exists, err := client.Exists(nil, key)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Printf("error checking existance: %s\n", err.Error())
			return
		}
		if !exists {
			log.Printf("error getting logged in user with id %s\n", id)
		} else {
			var u user
			err = client.GetObject(nil, key, &u)
			if err != nil {
				http.Error(res, "Server Error", http.StatusInternalServerError)
				log.Printf("error getting value: %s\n", err.Error())
				return
			}
			username = u.Username
		}
	}
	err = tpl.ExecuteTemplate(res, "index", username)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error running index template: %s\n", err.Error())
		return
	}
}

func loginPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "login", req.FormValue("msg"))
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error running login template: %s\n", err.Error())
		return
	}
}

func login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	name := req.FormValue("username")
	p := req.FormValue("password")

	var u *user

	// statements are query statements
	// namespace, set/kind, ...fieldsWeWant
	// note: usernames must be unique; not checked in this example
	stmt := as.NewStatement("bar", "users", "Username", "Password", "ID")
	stmt.Addfilter(as.NewEqualFilter("Username", name))
	rs, err := client.Query(nil, stmt)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error querying: %s\n", err.Error())
		return
	}
	defer rs.Close()
	for res := range rs.Results() {
		if res.Err != nil {
			log.Printf("error querying specific data: %s\n", err.Error())
		} else {
			u = &user{
				Username: res.Record.Bins["Username"].(string),
				Password: res.Record.Bins["Password"].(string),
				ID:       res.Record.Bins["ID"].(string),
			}
			break
		}
	}

	if u == nil {
		log.Printf("error logging in, no such user: %s\n", name)
		http.Redirect(res, req, "/login?msg=No such user", http.StatusSeeOther)
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) != nil {
		http.Redirect(res, req, "/login?msg=Incorrect password", http.StatusSeeOther)
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name:  "login",
		Value: u.ID,
	})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.SetCookie(res, &http.Cookie{
		Name:   "login",
		MaxAge: -1,
	})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func createPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "create", req.FormValue("msg"))
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error running create template: %s\n", err.Error())
		return
	}
}

func create(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	name := req.FormValue("username")
	p := req.FormValue("password")
	if len(name) < 3 || len(p) < 3 {
		http.Redirect(res, req, "/create?msg=Requires longer attributes", http.StatusSeeOther)
		return
	}
	id, err := uuid.NewV4()
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error generating uuid: %s\n", err.Error())
		return
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error hashing password: %s\n", err.Error())
		return
	}
	u := user{
		Username: name,
		Password: string(hashPass),
		ID:       id.String(),
	}

	key, err := as.NewKey("bar", "users", id.String())
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error creating key: %s\n", err.Error())
		return
	}
	err = client.PutObject(nil, key, &u)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Printf("error saving object: %s\n", err.Error())
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name:  "login",
		Value: id.String(),
	})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
