package movieinfo

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nu7hatch/gouuid"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
	"google.golang.org/appengine/urlfetch"
)

const githubAPIUrl = "https://api.github.com"

func renderMarkdown(ctx context.Context, text string) (string, error) {
	client := urlfetch.Client(ctx)
	response, err := client.Post(
		githubAPIUrl+"/markdown/raw",
		"text/x-markdown",
		strings.NewReader(text),
	)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	return string(bs), nil
}

type newMovieModel struct {
	CreatedID string
}

func handleNewMovie(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	model := &newMovieModel{}

	if req.Method == "POST" {
		title := req.FormValue("title")
		summary := req.FormValue("summary")
		poster, _, _ := req.FormFile("poster")

		var posterID string
		if poster != nil {
			defer poster.Close()
			pID, _ := uuid.NewV4()
			posterID = pID.String()
			err := putFile(ctx, posterID, poster)
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
		}

		summary, err := renderMarkdown(ctx, summary)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		index, err := search.Open("movies")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		movie := &Movie{
			Title:    title,
			Summary:  search.HTML(summary),
			PosterID: posterID,
		}

		id, err := index.Put(ctx, "", movie)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		model.CreatedID = id
	}

	err := tpl.ExecuteTemplate(res, "new-movie", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

/*
// with no checking to see if a poster was uploaded:

poster, _, err := req.FormFile("poster")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer poster.Close()

		posterID, _ := uuid.NewV4()

		err = putFile(ctx, posterID.String(), poster)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
*/
