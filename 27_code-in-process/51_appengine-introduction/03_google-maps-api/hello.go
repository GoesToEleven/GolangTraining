package main

import (
	"github.com/rwcarlsen/goexif/exif"
	"html/template"
	"log"
	"net/http"
	"os"
)

const GoogleAPIKey = "AIzaSyDpMNCWNz2UENVGQOS6zMFvtLsXn0zMBf4"

var tpls *template.Template

func init() {
	var err error
	tpls, err = template.ParseFiles("assets/templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {

	src, _ := os.Open("assets/img/IMG_20150714_191905.jpg")
	defer src.Close()
	x, _ := exif.Decode(src)
	lat, lon, _ := x.LatLong()

	var model struct {
		Latitude, Longitude float64
		Key                 string
	}
	model.Latitude = lat
	model.Longitude = lon
	model.Key = GoogleAPIKey
	err := tpls.ExecuteTemplate(res, "index.gohtml", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
