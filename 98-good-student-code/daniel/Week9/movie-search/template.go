package search

import "html/template"

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}
