package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type ToDo struct {
	ID    int64  `datastore:"-"`
	Email string `json:"-"`
	Text  string
}

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/todo", handleTodos)
	//	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	http.ServeFile(res, req, "assets/templates/index.html")
}

func handleTodos(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	todos := make([]ToDo, 0)

	switch req.Method {
	case "GET":
		q := datastore.NewQuery("ToDo").Filter("Email =", u.Email)
		iterator := q.Run(ctx)
		for {
			var todo ToDo
			key, err := iterator.Next(&todo)
			if err == datastore.Done {
				break
			} else if err != nil {
				log.Errorf(ctx, "error retrieving todos: %v", err)
				http.Error(res, err.Error(), 500)
				return
			}
			todo.ID = key.IntID()
			todos = append(todos, todo)
		}
		err := json.NewEncoder(res).Encode(todos)
		if err != nil {
			log.Errorf(ctx, "error marshalling todos: %v", err)
			return
		}
	case "POST":
		// the user is posting a new item
		var todo ToDo
		err := json.NewDecoder(req.Body).Decode(&todo)
		if err != nil {
			log.Errorf(ctx, "error unmarshalling todos: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		todo.Email = u.Email
		// add to datastore
		key := datastore.NewIncompleteKey(ctx, "ToDo", nil)
		key, err = datastore.Put(ctx, key, &todo)
		if err != nil {
			log.Errorf(ctx, "error adding todo: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		todo.ID = key.IntID()
		// send back to user
		err = json.NewEncoder(res).Encode(todo)
		if err != nil {
			log.Errorf(ctx, "error marshalling todo: %v", err)
			return
		}
	case "DELETE":
		id, _ := strconv.ParseInt(req.FormValue("id"), 10, 64)
		if id == 0 {
			http.Error(res, "not found", 404)
			return
		}
		key := datastore.NewKey(ctx, "ToDo", "", id, nil)
		var todo ToDo
		err := datastore.Get(ctx, key, &todo)
		if err != nil {
			log.Errorf(ctx, "error getting todo: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}
		if todo.Email != u.Email {
			http.Error(res, "access denied", 401)
			return
		}
		err = datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "error deleting todo: %v", err)
			return
		}
	default:
		http.Error(res, "Method Not Allowed", 405)
	}
}
