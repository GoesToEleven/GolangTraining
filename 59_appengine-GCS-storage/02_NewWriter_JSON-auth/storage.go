package storage

import (
	"io"
	"net/http"

	"golang.org/x/oauth2/google"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"io/ioutil"
	"golang.org/x/oauth2/jwt"
)

const gcsBucket = "learning-1130-bucket-01"
const aeId = "learning-1130"
var conf *jwt.Config

func init() {
	http.HandleFunc("/put", handlePut)

	jsonKey, err := ioutil.ReadFile("gcs.xxjson")
	if err != nil {
		panic("Couldn't read json key")
	}

	conf, err = google.JWTConfigFromJSON(
		jsonKey,
		storage.ScopeFullControl,
	)
	if err != nil {
		panic("Couldn't get *jwt.Config")
	}

}

func getCloudContext(req *http.Request) (context.Context, error) {
	ctx := appengine.NewContext(req)
	hc := conf.Client(ctx)
	return cloud.NewContext(aeId, hc), nil
}

func handlePut(res http.ResponseWriter, req *http.Request) {

	cctx, err := getCloudContext(req)
	if err != nil {
		http.Error(res, "ERROR GETTING CCTX: "+err.Error(), 500)
		return
	}

	writer := storage.NewWriter(cctx, gcsBucket, "myOffice.txt")
	writer.ContentType = "text/plain"
	io.WriteString(writer, "in my office")
	err = writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
}
