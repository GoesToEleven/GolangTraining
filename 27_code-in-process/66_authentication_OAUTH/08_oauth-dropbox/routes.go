package dropbox

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
)

func init() {
	r := httprouter.New()
	r.GET("/", handleIndex)
	r.GET("/login", handleLogin)
	r.GET("/oauth2", handleAuthorize)
	http.Handle("/", r)
}

func handleIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <a href="/login">Login with Dropbox</a>
  </body>
</html>`)
}

func handleLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	s := getSession(ctx, req)
	state, err := uuid.NewV4()
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	v := url.Values{}
	v.Add("response_type", "code")
	v.Add("client_id", "2be6c5b4z9uhar7")
	v.Add("redirect_uri", "http://localhost:8080/oauth2")
	v.Add("state", state.String())
	s.State = state.String()
	err = putSession(ctx, res, s)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	http.Redirect(res, req, "https://www.dropbox.com/1/oauth2/authorize?"+v.Encode(), http.StatusSeeOther)
}

type dropboxData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	UID         string `json:"uid"`
}

func handleAuthorize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	code, state := req.FormValue("code"), req.FormValue("state")
	s := getSession(ctx, req)
	if s.State != state {
		http.Error(res, "Detected cross-site attack", http.StatusUnauthorized)
		log.Criticalf(ctx, "Non-matching states from %s", req.RemoteAddr)
		return
	}
	data, err := getToken(ctx, code)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	io.WriteString(res, data.UID)
}

func getToken(ctx context.Context, code string) (*dropboxData, error) {
	v := url.Values{}
	v.Add("code", code)
	v.Add("grant_type", "authorization_code")
	v.Add("client_id", "2be6c5b4z9uhar7")
	v.Add("client_secret", "kdbp5bt12vodkz5")
	v.Add("redirect_uri", "http://localhost:8080/oauth2")
	client := urlfetch.Client(ctx)
	res, err := client.PostForm("https://api.dropbox.com/1/oauth2/token", v)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var data dropboxData
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	log.Debugf(ctx, "%v", data)
	return &data, nil
}
