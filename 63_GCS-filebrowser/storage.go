package filebrowser

import (
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

func getCloudContext(aeCtx context.Context) (context.Context, error) {
	data, err := ioutil.ReadFile("/Users/caleb/Desktop/gcs.json")
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(
		data,
		storage.ScopeFullControl,
	)
	if err != nil {
		return nil, err
	}

	tokenSource := conf.TokenSource(aeCtx)

	hc := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
			Base:   &urlfetch.Transport{Context: aeCtx},
		},
	}

	return cloud.NewContext(appengine.AppID(aeCtx), hc), nil
}

func listBucket(ctx context.Context, bucketName, folder string) ([]string, []string, error) {
	cloudContext, err := getCloudContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	var files, folders []string

	query := &storage.Query{
		Delimiter: "/",
		Prefix:    folder,
	}
	// objs is *storage.Objects
	objs, err := storage.ListObjects(cloudContext, bucketName, query)
	if err != nil {
		return nil, nil, err
	}

	for _, subfolder := range objs.Prefixes {
		folders = append(folders, subfolder[len(folder):])
	}

	for _, obj := range objs.Results {
		files = append(files, obj.Name)
	}

	return files, folders, nil
}

func putFile(ctx context.Context, bucketName, fileName string, rdr io.Reader) error {
	cctx, err := getCloudContext(ctx)
	if err != nil {
		return err
	}
	writer := storage.NewWriter(cctx, bucketName, fileName)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}
