package search

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
	"google.golang.org/appengine/urlfetch"
)

type movie struct {
	Name    string
	URL     string
	Summary search.HTML
}

func handleAdd(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	var err error

	if req.Method == "POST" {
		name := strings.TrimSpace(req.FormValue("name"))
		summary, err := getHTML(ctx, req.FormValue("summary"))
		if err != nil {
			http.Error(res, "Server error", http.StatusInternalServerError)
			log.Errorf(ctx, "%v\n", err)
			return
		}

		mov := &movie{
			Name:    name,
			Summary: summary,
			URL:     strings.ToLower(strings.Replace(name, " ", "", -1)),
		}
		err = addMovie(ctx, mov)
		if err != nil {
			http.Error(res, "Server error", http.StatusInternalServerError)
			log.Errorf(ctx, "%v\n", err)
			return
		}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	err = tpl.ExecuteTemplate(res, "addMovie", nil)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}
}

func getHTML(ctx context.Context, markdown string) (search.HTML, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Post("https://api.github.com/markdown/raw", "text/plain", strings.NewReader(markdown))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	return search.HTML(bts), err
}
