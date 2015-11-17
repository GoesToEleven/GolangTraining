package movieinfo

import (
	"io"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
	"io/ioutil"
	"google.golang.org/appengine/log"
)

const gcsBucket = "learning-1130-bucket-01"
const aeId = "learning-1130"

func getCloudContext(ctx context.Context) context.Context {
	jsonKey, err := ioutil.ReadFile("learning-860db08c451a.xxjson")
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil
	}

	conf, err := google.JWTConfigFromJSON(
		jsonKey,
		storage.ScopeFullControl,
	)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		return nil
	}

	hc := conf.Client(ctx)
	return cloud.NewContext(aeId, hc)
}

func putFile(ctx context.Context, name string, rdr io.Reader) error {
	cctx := getCloudContext(ctx)

	writer := storage.NewWriter(cctx, gcsBucket, name)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}

func getFile(ctx context.Context, name string) (io.ReadCloser, error) {
	cctx := getCloudContext(ctx)

	return storage.NewReader(cctx, gcsBucket, name)
}

func getFileLink(ctx context.Context, name string) (string, error) {
	cctx := getCloudContext(ctx)

	obj, err := storage.StatObject(cctx, gcsBucket, name)
	if err != nil {
		return "", err
	}
	return obj.MediaLink, nil
}
