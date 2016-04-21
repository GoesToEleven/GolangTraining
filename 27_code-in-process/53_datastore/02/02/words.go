package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

type Word struct {
	Term       string
	Definition string
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		term := req.FormValue("term")
		definition := req.FormValue("definition")

		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Word", term, 0, nil)

		entity := &Word{
			Term:       term,
			Definition: definition,
		}

		_, err := datastore.Put(ctx, key, entity)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
			<form method="POST" action="/words/">
				<input type="text" name="term">
				<textarea name="definition"></textarea>
				<input type="submit">
			</form>`)
}
