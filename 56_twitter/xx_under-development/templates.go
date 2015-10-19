package main

import (
	"html/template"
	"net/http"
	"github.com/dustin/go-humanize"
)

var tpl *template.Template

func init() {
	tpl = template.New("roottemplate")
	tpl = tpl.Funcs(template.FuncMap{
		"humanize_time": humanize.Time,
	})
	tpl = template.Must(tpl.ParseGlob("templates/html/*.html"))
}

func renderTemplate(res http.ResponseWriter, name string, data interface{}) {
	err := tpl.ExecuteTemplate(res, name, data)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}


// 	up above, you could have done it this way:
//		"humanize_time": func(tm time.Time) string {
//			return humanize.Time(tm)
//		},