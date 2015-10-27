package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}
	// the image doesn't serve
	io.WriteString(res, `
	Dog Name: <strong>`+dogName+`</strong><br>
	<img src="/toby.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(res, f)
}

func main() {
	http.HandleFunc("/toby.jpg", dogPic)
	http.HandleFunc("/dog/", upTown)
	http.ListenAndServe(":9000", nil)
}
