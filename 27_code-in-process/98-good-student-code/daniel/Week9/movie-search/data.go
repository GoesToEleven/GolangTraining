package search

import (
	"html/template"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

func getRecentMovies(ctx context.Context, maxItems int) ([]movie, error) {
	q := datastore.NewQuery("Movie").Project("Name", "URL").Limit(maxItems)
	var movies []movie
	_, err := q.GetAll(ctx, &movies)
	return movies, err
}

func addMovie(ctx context.Context, mov *movie) error {
	key := datastore.NewKey(ctx, "Movie", mov.URL, 0, nil)
	_, err := datastore.Put(ctx, key, mov)
	if err != nil {
		return err
	}
	return addToIndex(ctx, mov)
}

type templateMovie struct {
	Name    string
	URL     string
	Summary template.HTML
}

func getMovie(ctx context.Context, URL string) (*templateMovie, error) {
	key := datastore.NewKey(ctx, "Movie", URL, 0, nil)
	var mov templateMovie
	err := datastore.Get(ctx, key, &mov)
	return &mov, err
}
