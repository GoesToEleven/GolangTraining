package main

import "net/http"

func init() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/tweets", tweetHandler)
	http.HandleFunc("/api/follow", followHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
}
