package main

import (
	"bytes"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"io"
	"net/http"
)

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		var sd SessionData
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		tpl.ExecuteTemplate(res, templateName, sd)
	} else {
		ctx := appengine.NewContext(req)
		i, err := memcache.Get(ctx, templateName)
		if err != nil {
			buf := new(bytes.Buffer)
			writ := io.MultiWriter(res, buf)
			tpl.ExecuteTemplate(writ, templateName, SessionData{})
			memcache.Set(ctx, &memcache.Item{
				Key:   templateName,
				Value: buf.Bytes(),
			})
			return
		}
		io.WriteString(res, string(i.Value)) // we're serving the page from memcache
	}
}
