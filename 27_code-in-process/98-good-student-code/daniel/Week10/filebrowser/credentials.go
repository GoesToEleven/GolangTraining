package browser

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/julienschmidt/httprouter"
)

func handleCredentials(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "credentials", nil)
	if err != nil {
		ctx := appengine.NewContext(req)
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}

func handleLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	s := getSession(ctx, req)
	s.Bucket = req.FormValue("bucket")
	s.Credentials = req.FormValue("credentials")
	err := putSession(ctx, res, s)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	http.Redirect(res, req, "/browse/", http.StatusSeeOther)
}
