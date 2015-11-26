package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/getgif", handleGetGif)
}

func handleGetGif(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	result, err := client.Get("http://api.giphy.com/v1/gifs/search?q=funny+cat&api_key=dc6zaTOxFJmzC")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	defer result.Body.Close()
	var obj struct {
		Data []struct {
			URL    string `json:"url"`
			Images struct {
				Original struct {
					URL string
				}
			}
		}
	}
	err = json.NewDecoder(result.Body).Decode(&obj)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	for _, img := range obj.Data {
		fmt.Fprintf(res, `<a href="%v">%v</a><img src="%v"><br>`, img.URL, img.URL, img.Images.Original.URL)
	}
}

// https://github.com/Giphy/GiphyAPI
