package googleOauth2

import (
	"github.com/julienschmidt/httprouter"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"google.golang.org/api/gmail/v1"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/nu7hatch/gouuid"
	"io"
	"net/http"
)

var conf = &oauth2.Config{
	ClientID:     "979509136073-10ce3r5s8mka304l6od82t3nltp9cf8s.apps.googleusercontent.com",
	ClientSecret: "5zwpEL5WwwekeMsZoz6mcC0s",
	RedirectURL:  "https://practical-scion-114602.appspot.com/oauth2callback",
	Scopes:       []string{"https://www.googleapis.com/auth/gmail.readonly"},
	Endpoint:     google.Endpoint,
}

func init() {
	r := httprouter.New()
	r.GET("/", handleIndex)
	r.GET("/login", handleLogin)
	r.GET("/oauth2callback", handleAuthorize)
	http.Handle("/", r)
}

func handleIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
	<a href="/login">Login with Google to read your mail</a>
  </body>
</html>`)
}

func handleLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	//get session for storing state to check on callback
	s := getSession(ctx, req)

	//generate state for checking later
	state, err := uuid.NewV4()
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	//put state in session and putting to memcache to check on callback
	s.State = state.String()
	err = putSession(ctx, res, s)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	//generate the authorization url that goes to the login and consent page for google.
	//I set the ApprovalForce option so that it asks for consent each time so that we can see it.
	//Shows application asking to "Have offline access" each time. can remove second arg to remove this.
	url := conf.AuthCodeURL(s.State, oauth2.ApprovalForce)
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func handleAuthorize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	//retrieve code and state from the callback
	code, state := req.FormValue("code"), req.FormValue("state")

	//get session from memcache
	s := getSession(ctx, req)

	//compare state from callback with state stored in memcache
	if state != s.State {
		http.Error(res, "Detected cross-site attack", http.StatusUnauthorized)
		log.Criticalf(ctx, "Non-matching states from %s", req.RemoteAddr)
		return
	}

	//exchange the auth code given for an access token
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	//create a client from the token
	client := conf.Client(ctx, tok)

	//create a gmail service from the client
	srv, err := gmail.New(client)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	//request a list of messages in the user's mailbox
	list, err := srv.Users.Threads.List("me").Do()
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	//loop through and print out the first 25 message threads from your mailbox
	output := "<p>Messages:</p>"
	if len(list.Threads) > 0 {
		i := 1
		for _, val := range list.Threads {
			output += "<p>- " + val.Snippet + "...</p>"
			if i > 25 {
				break
			}
			i++
		}
	} else {
		output += "<p>You have no messages!</p>"
	}
	io.WriteString(res, output)
}
