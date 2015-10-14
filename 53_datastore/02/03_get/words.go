package main

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", handleWords)
}

type Word struct {
	Term       string
	Definition string
}

func handleWords(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		term := strings.Split(req.URL.Path, "/")[1]
		showWord(res, req, term)
		return
	}

	if req.Method == "POST" {
		saveWord(res, req)
		return
	}

	listWords(res, req)
}

func listWords(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Word").Order("Term")

	html := ""

	iterator := q.Run(ctx)
	for {
		var entity Word
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		html += `
			<dt>` + entity.Term + `</dt>
			<dd>` + entity.Definition + `</dd>
		`
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
			<dl>
				`+html+`
			</dl>
			<form method="POST">
				<input type="text" name="term">
				<textarea name="definition"></textarea>
				<input type="submit">
			</form>
			`)
}

func showWord(res http.ResponseWriter, req *http.Request, term string) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Word", term, 0, nil)
	var entity Word
	err := datastore.Get(ctx, key, &entity)
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<dl>
			<dt>`+entity.Term+`</dt>
			<dd>`+entity.Definition+`</dd>
		</dl>
	`)
}

func saveWord(res http.ResponseWriter, req *http.Request) {
	term := req.FormValue("term")
	definition := req.FormValue("definition")
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Word", term, 0, nil)
	entity := Word{
		Term:       term,
		Definition: definition,
	}

	_, err := datastore.Put(ctx, key, &entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	http.Redirect(res, req, "/", 302)
}
