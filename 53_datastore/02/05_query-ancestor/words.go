package main

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"log"
)

func init() {
	http.HandleFunc("/", handleAnimals)
	http.HandleFunc("/ateprocess", ateProcess)
}

type Animal struct {
	Species     string
	Description string
}

type Ate struct {
	FoodItem string
}

func handleAnimals(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		term := strings.Split(req.URL.Path, "/")[1]
		showAnimal(res, req, term)
		return
	}
	if req.Method == "POST" {
		saveAnimal(res, req)
		return
	}
	listAnimals(res, req)
}

func listAnimals(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Animal").Order("Species")

	html := ""

	iterator := q.Run(ctx)
	for {
		var entity Animal
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		html += `
			<dt>` + entity.Species + `</dt>
			<dd>` + entity.Description + `</dd>
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

func showAnimal(res http.ResponseWriter, req *http.Request, term string) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Animal", term, 0, nil)
	var entity Animal
	err := datastore.Get(ctx, key, &entity)
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	// have they eaten something?
	q := datastore.NewQuery("Ate").Ancestor(key)
	html := ""
	iterator := q.Run(ctx)
	for {
		var entity Ate
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		html += `
			<li>` + entity.FoodItem + `</li>
		`
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<dl>
			<dt>`+entity.Species+`</dt>
			<dd>`+entity.Description+`</dd>
		</dl>
		<h1>This Animal Just Ate:</h1>
		<form method="POST" action="ateprocess">
			<textarea name="fooditem"></textarea>
			<input type="hidden" name="eater" value="`+term+`">
			<input type="submit">
		</form>
		<h1>This Animal Has Already Eaten:</h1>
		<ol>`+html+`</ol>
	`)
}

func saveAnimal(res http.ResponseWriter, req *http.Request) {
	term := req.FormValue("term")
	definition := req.FormValue("definition")
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Animal", term, 0, nil)
	entity := Animal{
		Species:     term,
		Description: definition,
	}

	_, err := datastore.Put(ctx, key, &entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	http.Redirect(res, req, "/", 302)
}

func ateProcess(res http.ResponseWriter, req *http.Request) {
	fooditem := req.FormValue("fooditem")
	eater := req.FormValue("eater")
	ctx := appengine.NewContext(req)
	eaterKey := datastore.NewKey(ctx, "Animal", eater, 0, nil)
	key := datastore.NewIncompleteKey(ctx, "Ate", eaterKey)
	entity := Ate{
		FoodItem: fooditem,
	}
	log.Println(fooditem)
	log.Println(entity)
	log.Println(&entity)
	_, err := datastore.Put(ctx, key, &entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	http.Redirect(res, req, "/", 302)
}
