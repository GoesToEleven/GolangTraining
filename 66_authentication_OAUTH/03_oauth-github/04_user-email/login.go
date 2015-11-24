package githubexample

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
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

	fmt.Fprintln(res, "AUTHORIZATION CODE "+code)

	accessToken, err := getAccessToken(ctx, state, code)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	fmt.Fprintln(res, "ACCESS TOKEN "+accessToken)

	email, err := getEmail(ctx, accessToken)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	fmt.Fprintln(res, "EMAIL "+email)

}

func getAccessToken(ctx context.Context, state, code string) (string, error) {
	values := make(url.Values)
	values.Add("client_id", "fbbaa8ce5c394b7c3198")
	values.Add("client_secret", "1b450ffb26982847d1c92eadd8a6d4932a79f225")
	values.Add("code", code)
	values.Add("state", state)
	client := urlfetch.Client(ctx)
	response, err := client.PostForm("https://github.com/login/oauth/access_token", values)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	bs, _ := ioutil.ReadAll(response.Body)
	v, _ := url.ParseQuery(string(bs))
	return v.Get("access_token"), nil
}

func getEmail(ctx context.Context, accessToken string) (string, error) {
	client := urlfetch.Client(ctx)
	response, err := client.Get("https://api.github.com/user/emails?access_token=" + accessToken)
	if err != nil {
		return "", nil
	}
	defer response.Body.Close()

	var data []struct {
		Email    string
		Verified bool
		Primary  bool
	}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", nil
	}
	if len(data) == 0 {
		return "", fmt.Errorf("no email found")
	}
	return data[0].Email, nil
}
