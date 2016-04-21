package example

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type Entity struct {
	Value string
}

func init() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func home(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	k := datastore.NewKey(c, "Entity", "todd.mcleod@fresnocitycollege.edu", 0, nil)
	e := new(Entity)
	if err := datastore.Get(c, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	old := e.Value
	e.Value = r.URL.Path

	if _, err := datastore.Put(c, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "old=%q\nnew=%q\n", old, e.Value)
}
