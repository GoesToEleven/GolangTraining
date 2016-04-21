package movieinfo

import "google.golang.org/appengine/search"

type Movie struct {
	Title    string
	Summary  search.HTML
	PosterID string
}
