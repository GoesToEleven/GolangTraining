package main

import (
	"bytes"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"io"
	"net/http"
)

func memTemplate(res http.ResponseWriter, req *http.Request, memKey, templateName string) {
	ctx := appengine.NewContext(req)
	i, err := memcache.Get(ctx, memKey)
	if err != nil {
		buf := new(bytes.Buffer)
		writ := io.MultiWriter(res, buf)
		tpl.ExecuteTemplate(writ, templateName, nil)
		memcache.Set(ctx, &memcache.Item{
			Value: buf.Bytes(),
			Key:   memKey,
		})
		return
	}
	io.WriteString(res, string(i.Value))
}
