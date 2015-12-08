package githubexample

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"github.com/nu7hatch/gouuid"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/github-login", handleGithubLogin)
	http.HandleFunc("/oauth2callback", handleOauth2Callback)
	http.HandleFunc("/github-info", handleGithubInfo)
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
	values.Add("client_id", "767154e6915134caade5")
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
	api := NewGithubAPI(ctx)
	// get the session
	session := getSession(ctx, req)

	state := req.FormValue("state")
	code := req.FormValue("code")

	if state != session.State {
		http.Error(res, "invalid state", 401)
		return
	}

	accessToken, err := api.getAccessToken(state, code)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	api.accessToken = accessToken

	username, err := api.getUsername()
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	session.Username = username
	session.AccessToken = accessToken
	putSession(ctx, res, session)

	delayedGetStats.Call(ctx, accessToken, username)
	http.Redirect(res, req, "/github-info", 302)

}

func handleGithubInfo(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	session := getSession(ctx, req)

	var stats CommitStats

	key := datastore.NewKey(ctx, "Stats", session.Username, 0, nil)
	err := datastore.Get(ctx, key, &stats)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	// stats, err := api.getCommitSummaryStats(since)
	// if err != nil {
	// 	http.Error(res, err.Error(), 500)
	// 	return
	// }

	io.WriteString(res, `<!DOCTYPE html>
<html>
	<head>

	</head>
	<body>
		In the last month you have:
		<table>
			<tr>
				<th>additions</th>
				<td>`+fmt.Sprint(stats.Additions)+`</td>
			</tr>
			<tr>
				<th>deletions</th>
				<td>`+fmt.Sprint(stats.Deletions)+`</td>
			</tr>
		</table>
	</body>
</html>`)
}
