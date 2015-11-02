package twitter

import (
	"html/template"
	"io"
	"net/http"
	"net/mail"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type profile struct {
	Username  string
	Email     string
	Following []string
}

type tweet struct {
	Username   string `datastore:"-"`
	Message    string
	SubmitTime time.Time
}

type mainpageData struct {
	Tweets   []tweet
	Logged   bool
	Username string
}

type profileData struct {
	Tweets      []tweet
	Profile     profile
	CurrentUser string
	Following   bool
}

type loginData struct {
	ErrorMessage string
	Username     string
}

const (
	minUsernameSize = 5
	maxUsernameSize = 20
	tweetSize       = 140
	loginDuration   = 60 * 60 * 24 // 1 Day
)

var tpl = template.New("templates")

func init() {
	_, err := tpl.ParseGlob("templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/CreateProfile", handleCreateProfile)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/tweet.json", handleTweet)
	http.HandleFunc("/_ah/mail/", incomingMail)
}

func getCurrentUser(req *http.Request) *profile {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	if u == nil {
		return nil
	}
	c, err := req.Cookie("login")
	if err != nil {
		return nil
	}
	p, err := getProfileByUsername(ctx, c.Value)
	if err != nil {
		return nil
	}
	if p.Email != u.Email {
		return nil
	}
	return p
}

func confirmCreateProfile(ctx context.Context, username string) bool {
	_, err := getProfileByUsername(ctx, username)
	return len(username) >= minUsernameSize && len(username) <= maxUsernameSize &&
		err == datastore.ErrNoSuchEntity
}

func incomingMail(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	defer req.Body.Close()
	msg, err := mail.ReadMessage(req.Body)
	if err != nil {
		log.Errorf(ctx, "Email error: %s\n", err.Error())
		return
	}
	addresses, err := msg.Header.AddressList("From")
	if err != nil {
		log.Errorf(ctx, "Email error: %s\n", err.Error())
		return
	}
	addr := addresses[0]
	p, err := getProfileByEmail(ctx, addr.Address)
	if err != nil {
		log.Errorf(ctx, "Email error: %s\n", err.Error())
		return
	}

	contentType := msg.Header.Get("Content-Type")
	text, err := parseFile(ctx, contentType, "", msg.Body)
	if err != nil {
		log.Errorf(ctx, "Email error: %s\n", err.Error())
		return
	}
	t := &tweet{
		Message:    text,
		SubmitTime: time.Now(),
		Username:   p.Username,
	}
	err = postTweet(ctx, t, addr.Address)
	if err != nil {
		log.Errorf(ctx, "Email error: %s\n", err.Error())
		return
	}
}

func handleTweet(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := getCurrentUser(req)
	if u == nil {
		http.Error(res, "Incorrect login", http.StatusUnauthorized)
		log.Warningf(ctx, "Incorrect login from: %s\n", req.RemoteAddr)
		return
	}
	if req.Method != "POST" {
		http.Error(res, "Unknown method", http.StatusMethodNotAllowed)
		log.Warningf(ctx, "Incorrect method on tweet.json from %s", req.RemoteAddr)
		return
	}
	buffer := make([]byte, tweetSize)
	n, err := req.Body.Read(buffer)
	if err != nil && err != io.EOF {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		log.Warningf(ctx, "Bad request: %s\n", err.Error())
		return
	}
	msg := string(buffer[:n])
	t := tweet{
		Username:   u.Username,
		Message:    msg,
		SubmitTime: time.Now(),
	}
	err = postTweet(ctx, &t, u.Email)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Put Tweet Error: %s\n", err.Error())
		return
	}
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	cookie, err := req.Cookie("login")
	if err != http.ErrNoCookie {
		http.Redirect(res, req, "/"+cookie.Value, http.StatusSeeOther)
		return
	}

	currentProfile, err := getProfileByEmail(ctx, u.Email)
	if err == datastore.ErrNoSuchEntity {
		http.Redirect(res, req, "/CreateProfile", http.StatusSeeOther)
		return
	} else if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Get profile error: %s\n", err.Error())
		return
	}

	login := loginData{
		Username: currentProfile.Username,
	}
	if req.Method == "POST" {
		username := req.FormValue("username")
		p, err := getProfileByUsername(ctx, username)
		if err != nil {
			login.ErrorMessage = "No such username"
		} else if p.Email != u.Email {
			login.ErrorMessage = "Not your profile"
		} else {
			c := http.Cookie{
				Name:   "login",
				Value:  username,
				MaxAge: loginDuration,
			}
			http.SetCookie(res, &c)
			http.Redirect(res, req, "/"+username, http.StatusSeeOther)
			return
		}
	}
	err = tpl.ExecuteTemplate(res, "login.gohtml", login)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

func handleLogout(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{Name: "login", MaxAge: -1})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func handleCreateProfile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	if req.Method == "POST" {
		username := req.FormValue("username")
		if !confirmCreateProfile(ctx, username) {
			http.Error(res, "Invalid input!", http.StatusBadRequest)
			log.Warningf(ctx, "Invalid profile information from %s\n", req.RemoteAddr)
			return
		}
		err := createProfile(ctx, username, u.Email)
		http.SetCookie(res, &http.Cookie{Name: "login", Value: username, MaxAge: loginDuration})
		if err != nil {
			http.Error(res, "Server error!", http.StatusInternalServerError)
			log.Errorf(ctx, "Create profile Error: %s\n", err.Error())
			return
		}
	}

	_, err := getProfileByEmail(ctx, u.Email)
	if err == nil {
		http.Redirect(res, req, "login", http.StatusSeeOther)
		return
	} else if err != datastore.ErrNoSuchEntity {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Get profile Error: %s\n", err.Error())
		return
	}

	err = tpl.ExecuteTemplate(res, "createProfile.gohtml", nil)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

func getProfile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	username := req.URL.Path[1:]

	p, err := getProfileByUsername(ctx, username)
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Get Profile Error: %s\n", username)
		return
	}

	tweets, err := getTweets(ctx, p.Email)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Query Error: %s\n", err.Error())
		return
	}

	pd := profileData{
		Tweets:  tweets,
		Profile: *p,
	}
	u := getCurrentUser(req)
	if u != nil {
		pd.CurrentUser = u.Username
		pd.Following = itemIn(username, u.Following)
	}

	isFollower := req.URL.Query().Get("f")
	if isFollower == "y" {
		err := addFollower(ctx, u, username)
		if err != nil {
			http.Error(res, "Unable to remove follower", http.StatusInternalServerError)
			log.Errorf(ctx, "Save add followed: %s\n", err.Error())
			return
		}
	} else if isFollower == "n" {
		err := removeFollower(ctx, u, username)
		if err != nil {
			http.Error(res, "Unable to remove follower", http.StatusInternalServerError)
			log.Errorf(ctx, "Save remove followed: %s\n", err.Error())
			return
		}
	}

	err = tpl.ExecuteTemplate(res, "profile.gohtml", pd)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}

func handle(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		getProfile(res, req)
		return
	}

	ctx := appengine.NewContext(req)
	u := getCurrentUser(req)

	if u != nil {
		_, err := getProfileByEmail(ctx, u.Email)
		if err == datastore.ErrNoSuchEntity {
			http.Redirect(res, req, "/CreateProfile", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(res, "Server error!", http.StatusInternalServerError)
			log.Errorf(ctx, "Datastore get Error: %s\n", err.Error())
			return
		}
	}

	// Get recent tweets
	var tweets []tweet
	var err error
	if u == nil {
		tweets, err = getTweets(ctx, "")
	} else {
		tweets, err = getMultiTweets(ctx, u.Following)
	}

	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Query Error: %s\n", err.Error())
		return
	}

	// Create template
	data := mainpageData{
		Tweets: tweets,
	}

	c, err := req.Cookie("login")
	if err == nil {
		data.Logged = true
		data.Username = c.Value
	} else {
		data.Logged = false
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		http.Error(res, "Server error!", http.StatusInternalServerError)
		log.Errorf(ctx, "Template Execute Error: %s\n", err.Error())
		return
	}
}
