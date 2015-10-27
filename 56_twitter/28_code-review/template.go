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
	session := getSession(req)
	if len(session.Value) > 0 {
		var sd sessionData
		json.Unmarshal(session.Value, &sd)
		sd.LoggedIn = true
		tpl.ExecuteTemplate(res, templateName, sd)
	} else {
		ctx := appengine.NewContext(req)
		i, err := memcache.Get(ctx, templateName)
		if err != nil {
			buf := new(bytes.Buffer)
			writ := io.MultiWriter(res, buf)
			tpl.ExecuteTemplate(writ, templateName, sessionData{})
			memcache.Set(ctx, &memcache.Item{
				Value: buf.Bytes(),
				Key:   templateName,
			})
			return
		}
		io.WriteString(res, string(i.Value)) // we're serving the page from memcache
	}
}
