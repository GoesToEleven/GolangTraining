package search

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

func handleSearch(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	movies, err := searchMovies(ctx, req.FormValue("q"))
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}

	err = tpl.ExecuteTemplate(res, "search", movies)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, "%v\n", err)
		return
	}
}

func addToIndex(ctx context.Context, mov *movie) error {
	idx, err := search.Open("movies")
	if err != nil {
		return err
	}

	_, err = idx.Put(ctx, "", mov)
	return err
}

func searchMovies(ctx context.Context, query string) ([]movie, error) {
	idx, err := search.Open("movies")
	if err != nil {
		return nil, err
	}

	it := idx.Search(ctx, query, nil)
	movies := []movie{}
	for {
		var mov movie
		_, err := it.Next(&mov)
		if err == search.Done {
			break
		} else if err != nil {
			return nil, err
		}
		movies = append(movies, mov)
	}
	return movies, nil
}
