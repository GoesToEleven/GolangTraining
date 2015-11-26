package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if len(ps) != 0 {
		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
		return
	}

	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/user/*name", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
