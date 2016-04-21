package main

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Employee struct {
	Name     string
	Role     string
	HireDate time.Time
	Account  string
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	e1 := Employee{
		Name:     "Joe Citizen",
		Role:     "Manager",
		HireDate: time.Now(),
		Account:  user.Current(c).String(),
	}

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "employee", nil), &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 Employee
	if err = datastore.Get(c, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored and retrieved the Employee named %q", e2.Name)
}
