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
)

const gcsBucket = "learning-1130-bucket-01"
const aeId = "learning-1130"

func init() {
	http.HandleFunc("/put", handlePut)
}

func getCloudContext(req *http.Request) (context.Context, error) {
	jsonKey, err := ioutil.ReadFile("learning-860db08c451a.Xjson")
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

	writer := storage.NewWriter(cctx, gcsBucket, "exampleJSON2.txt")
	io.WriteString(writer, "AGAIN WITH JSON AUTH")
	err = writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
}
