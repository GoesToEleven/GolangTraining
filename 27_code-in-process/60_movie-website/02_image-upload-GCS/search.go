package movieinfo

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
)

type searchModel struct {
	Query  string
	Movies []Movie
}

func handleSearch(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	query := req.FormValue("q")
	model := &searchModel{
		Query: query,
	}

	index, err := search.Open("movies")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	iterator := index.Search(ctx, query, nil)
	for {
		var movie Movie
		_, err := iterator.Next(&movie)
		if err == search.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		model.Movies = append(model.Movies, movie)
	}

	err = tpl.ExecuteTemplate(res, "search", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}
