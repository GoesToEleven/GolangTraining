package movieinfo

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
)

type newMovieModel struct {
	CreatedID string
}

func handleNewMovie(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	model := &newMovieModel{}

	if req.Method == "POST" {
		title := req.FormValue("title")
		summary := req.FormValue("summary")

		index, err := search.Open("movies")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		movie := &Movie{
			Title:   title,
			Summary: summary,
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
