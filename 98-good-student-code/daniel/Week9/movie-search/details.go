package search

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func handleMovieDetails(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	mov, err := getMovie(ctx, req.URL.Path[1:])
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}
	err = tpl.ExecuteTemplate(res, "details", mov)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}
}
