package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/todo.json", jsonServe)
}

func handle(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		http.ServeFile(res, req, "index.html")
	case "/style.css":
		http.ServeFile(res, req, "style.css")
	case "/script.js":
		http.ServeFile(res, req, "script.js")
	default:
		http.NotFound(res, req)
	}
}

type dataItem struct {
	Value  string
	KeyVal string
}

type list struct {
	Value string
	Owner string
	Time  time.Time
}

func jsonServe(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if len(req.FormValue("keyVal")) > 0 {
			deleteData(res, req)
		} else {
			saveJSON(res, req)
		}
	case "GET":
		getJSON(res, req)
	default:
		ctx := appengine.NewContext(req)
		http.Error(res, "Bad HTTP method", http.StatusBadRequest)
		log.Warningf(ctx, "Attempted HTTP method %s from %s", req.Method, req.RemoteAddr)
	}
}

func deleteData(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	keyVal := req.FormValue("keyVal")
	key, err := datastore.DecodeKey(keyVal)
	if err != nil {
		http.Error(res, "Invalid data", http.StatusBadRequest)
		log.Warningf(ctx, err.Error())
		return
	}
	var l list
	err = datastore.Get(ctx, key, &l)
	if err != nil {
		http.Error(res, "Invalid data", http.StatusBadRequest)
		log.Warningf(ctx, err.Error())
		return
	}
	if l.Owner != u.Email {
		http.Error(res, "Not authorized to delete this entry", http.StatusUnauthorized)
		log.Warningf(ctx, err.Error())
		return
	}
	err = datastore.Delete(ctx, key)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}

func getQuery(email string) *datastore.Query {
	return datastore.NewQuery("List").
		Filter("Owner =", email).
		Order("Time")
}

func saveJSON(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	decoder := json.NewDecoder(req.Body)
	key := datastore.NewIncompleteKey(ctx, "List", nil)
	var l list
	err := decoder.Decode(&l)
	if err != nil {
		http.Error(res, "Invalid data", http.StatusBadRequest)
		log.Warningf(ctx, err.Error())
		return
	}
	l.Owner = u.Email
	l.Time = time.Now()
	_, err = datastore.Put(ctx, key, &l)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}

func getJSON(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	query := getQuery(u.Email)
	l := []list{}
	keys, err := query.GetAll(ctx, &l)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	items := make([]dataItem, len(l))
	for i := range l {
		items[i].Value = l[i].Value
		items[i].KeyVal = keys[i].Encode()
	}
	enc := json.NewEncoder(res)
	err = enc.Encode(items)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}
