package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Printf("%T %v \n", ps, ps)
	fmt.Println(ps == nil)    // true
	fmt.Println(ps != nil)    // false
	fmt.Println(len(ps) == 0) // true
	fmt.Println(len(ps) != 0) // false
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/:name", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
