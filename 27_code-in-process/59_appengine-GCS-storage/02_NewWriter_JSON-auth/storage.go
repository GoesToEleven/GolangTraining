package storage

import (
	"io"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/cloud/storage"
)

const gcsBucket = "learning-1130-bucket-01"

func init() {
	http.HandleFunc("/put", handlePut)
}

func handlePut(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		http.Error(res, "ERROR CREATING NEW CLIENT: "+err.Error(), 500)
		return
	}
	writer := client.Bucket(gcsBucket).Object("myOffice.txt").NewWriter(ctx)
	writer.ContentType = "text/plain"
	io.WriteString(writer, "in my office")
	err = writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
	client.Close()
}
