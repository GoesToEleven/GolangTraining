package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
	ID       string
}

var tpl *template.Template
var users = map[string]user{}
var idUsers = map[string]user{}

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	f, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	loadUsers()

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/login", loginPage)
	router.POST("/login", login)
	router.GET("/logout", logout)
	router.GET("/create", createPage)
	router.POST("/create", create)
	go func() {
		err = http.ListenAndServe(":9000", router)
		if err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	saveUsers()
}

func loadUsers() {
	var rdr io.Reader
	f, err := os.Open("users.json")
	if err != nil {
		rdr = strings.NewReader("{}")
	} else {
		defer f.Close()
		rdr = f
	}
	err = json.NewDecoder(rdr).Decode(&users)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		idUsers[u.ID] = u
	}
}

func saveUsers() {
	f, err := os.Create("users.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(users)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var username string
	c, err := req.Cookie("login")
	if err == nil {
		id := c.Value
		u, ok := idUsers[id]
		if !ok {
			log.Printf("error getting logged in user with id %s\n", id)
		} else {
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

	u, ok := users[name]
	if !ok {
		log.Printf("error logging in, no such user: %s\n", name)
		http.Redirect(res, req, "/login?msg=No such user", http.StatusSeeOther)
		return
	}
	if bcrypt.CompareHashAndPassword(u.Password, []byte(p)) != nil {
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
		Password: hashPass,
		ID:       id.String(),
	}
	users[name] = u
	idUsers[id.String()] = u
	http.SetCookie(res, &http.Cookie{
		Name:  "login",
		Value: id.String(),
	})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
