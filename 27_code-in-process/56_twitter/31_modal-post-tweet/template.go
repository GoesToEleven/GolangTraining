package main

import (
	"encoding/json"
	"net/http"
)

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem := getSession(req)
	if len(memItem.Value) > 0 {
		// logged in
		var sd SessionData
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		tpl.ExecuteTemplate(res, templateName, sd)
	} else {
		// not logged in
		tpl.ExecuteTemplate(res, templateName, SessionData{})
	}
}
