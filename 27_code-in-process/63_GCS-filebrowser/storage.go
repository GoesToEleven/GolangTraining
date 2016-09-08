package filebrowser

import (
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
	"io"
)

func listBucket(ctx context.Context, bucketName, folder string) ([]string, []string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	var files, folders []string

	query := &storage.Query{
		Delimiter: "/",
		Prefix:    folder,
	}
	// objs is *storage.Objects
	objs, err := client.Bucket(bucketName).List(ctx, query)
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
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	writer := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}
