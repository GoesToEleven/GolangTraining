package githubexample

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
)

// change redirectURI for deployment; eg, http://<yourAppId>.appspot.com/oauth2callback
const redirectURI = "http://localhost:8080/oauth2callback"
const githubAPIURL = "https://api.github.com"

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/github-login", handleGithubLogin)
	http.HandleFunc("/oauth2callback", handleOauth2Callback)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <a href="/github-login">LOGIN WITH GITHUB</a>
  </body>
</html>`)
}

var githubScopes = []string{
	"user:email",
	"read:org",
}

func handleGithubLogin(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// get the session
	session := getSession(ctx, req)
	id, _ := uuid.NewV4()

	values := make(url.Values)
	values.Add("client_id", "fbbaa8ce5c394b7c3198")
	values.Add("redirect_uri", redirectURI)
	values.Add("scope", strings.Join(githubScopes, ","))
	values.Add("state", id.String())

	// save the session
	session.State = id.String()
	putSession(ctx, res, session)

	http.Redirect(res, req, fmt.Sprintf(
		"https://github.com/login/oauth/authorize?%s",
		values.Encode(),
	), 302)
}

func handleOauth2Callback(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// get the session
	session := getSession(ctx, req)

	state := req.FormValue("state")
	code := req.FormValue("code")

	if state != session.State {
		http.Error(res, "invalid state", 401)
		return
	}

	fmt.Fprintln(res, code)
}
