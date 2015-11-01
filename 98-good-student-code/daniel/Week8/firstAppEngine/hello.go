package hello

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/user"
)

func init() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	globalCounter, err := memcache.Increment(ctx, "globalCounter", 1, 0)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
	}
	currentUser := user.Current(ctx)
	userCounter, err := memcache.Increment(ctx, currentUser.Email+"-Counter", 1, 0)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
	}
	fmt.Fprintf(res, "Hello World! You are visitor #%d, and you have visited %d times!", globalCounter, userCounter)
}
