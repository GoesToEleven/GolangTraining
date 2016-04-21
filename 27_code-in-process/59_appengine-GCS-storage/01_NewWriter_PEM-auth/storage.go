package storage

import (
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

func init() {
	http.HandleFunc("/put", handlePut)

}

func handlePut(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	bucket := "learning-1130-bucket-01"

	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(ctx, storage.ScopeFullControl),
			Base:   &urlfetch.Transport{Context: ctx},
		},
	}

	cctx := cloud.NewContext(appengine.AppID(ctx), hc)

	writer := storage.NewWriter(cctx, bucket, "example.txt")
	io.WriteString(writer, "Hello World!!!!!")
	err := writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
}
