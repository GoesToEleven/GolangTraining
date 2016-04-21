package search

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		handleMovieDetails(res, req)
		return
	}

	ctx := appengine.NewContext(req)
	data := struct {
		Movies []movie
	}{}
	m, err := getRecentMovies(ctx, 15)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}
	data.Movies = m

	err = tpl.ExecuteTemplate(res, "index", data)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}
}
