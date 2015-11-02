package twitter

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"strings"

	"golang.org/x/net/context"
)

func parseFile(ctx context.Context, mme, enc string, reader io.Reader) (string, error) {
	mediatype, params, err := mime.ParseMediaType(mme)
	ret := ""
	if err != nil {
		return "", err
	}
	switch mediatype {
	case "multipart/alternative":
		rdr := multipart.NewReader(reader, params["boundary"])
		for {
			part, err := rdr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				return "", err
			}
			newMime := part.Header.Get("Content-Type")
			if !strings.HasPrefix(newMime, "text/plain") {
				continue
			}
			encoding := part.Header.Get("Content-Transfer-Encoding")
			ret, err = parseFile(ctx, newMime, encoding, part)
			if err != nil {
				return "", err
			}
		}

	case "text/plain":
		if enc == "base64" {
			dcode := base64.NewDecoder(base64.StdEncoding, reader)
			bts, err := ioutil.ReadAll(dcode)
			if err != nil {
				return "", err
			}
			ret = string(bts)
		} else {
			bts, err := ioutil.ReadAll(reader)
			if err != nil {
				return "", err
			}
			ret = string(bts)
		}
	}

	return ret, nil
}
