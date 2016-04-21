package storage

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/appengine"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"io/ioutil"
)

const gcsBucket = "learning-1130-bucket-01"
const aeId = "learning-1130"

func init() {
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/list", handleList)
}

func getCloudContext(req *http.Request) (context.Context, error) {
	jsonKey, err := ioutil.ReadFile("gcs.xxjson")
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(
		jsonKey,
		storage.ScopeFullControl,
	)
	if err != nil {
		return nil, err
	}

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

	writer := storage.NewWriter(cctx, gcsBucket, "example999.txt")
	io.WriteString(writer, "showing the put get list")
	err = writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	cctx, err := getCloudContext(req)
	if err != nil {
		http.Error(res, "ERROR GETTING CCTX: "+err.Error(), 500)
		return
	}

	rdr, err := storage.NewReader(cctx, gcsBucket, "example999.txt")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	defer rdr.Close()

	io.Copy(res, rdr)
}

func handleList(res http.ResponseWriter, req *http.Request) {
	cctx, err := getCloudContext(req)
	if err != nil {
		http.Error(res, "ERROR GETTING CCTX: "+err.Error(), 500)
		return
	}

	var query *storage.Query
	objs, err := storage.ListObjects(cctx, gcsBucket, query)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	for _, obj := range objs.Results {
		fmt.Fprintln(res, obj.Name)
	}
}
