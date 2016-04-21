package storage

import (
	"html/template"
	"io"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

const (
	bucket = "golang-bootcamp"
	prefix = "Oralordos/"
)

func init() {
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/", handleList)
}

func getCloudContext(ctx context.Context) context.Context {
	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.AppEngineTokenSource(ctx, storage.ScopeFullControl),
			Base:   &urlfetch.Transport{Context: ctx},
		},
	}

	return cloud.NewContext(appengine.AppID(ctx), hc)
}

func handlePut(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	if req.Method != "POST" {
		http.Error(res, "Must send a file", http.StatusMethodNotAllowed)
		return
	}

	file, hdr, err := req.FormFile("f")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	filename := hdr.Filename

	writer := storage.NewWriter(cctx, bucket, prefix+filename)
	io.Copy(writer, file)
	err = writer.Close()
	if err != nil {
		http.Error(res, "ERROR WRITING TO BUCKET: "+err.Error(), 500)
		return
	}
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	filename := req.FormValue("f")

	rdr, err := storage.NewReader(cctx, bucket, prefix+filename)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	defer rdr.Close()

	io.Copy(res, rdr)
}

func handleList(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	cctx := getCloudContext(ctx)

	query := &storage.Query{
		Prefix: prefix,
	}
	objs, err := storage.ListObjects(cctx, bucket, query)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	t, err := template.New("").Parse(`<li><a href="/get?f={{.}}">{{.}}</a></li>`)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, `<form action="put" method="POST" enctype="multipart/form-data"><input name="f" type="file"><input type="submit"></form> <ul>`)
	for _, obj := range objs.Results {
		err := t.Execute(res, obj.Name[len(prefix):])
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}
	io.WriteString(res, `</ul>`)
}
