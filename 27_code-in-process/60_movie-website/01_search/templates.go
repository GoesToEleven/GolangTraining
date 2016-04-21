package movieinfo

import "github.com/alecthomas/template"

var tpl = template.Must(
	template.New("").
		Funcs(template.FuncMap{}).
		ParseGlob("templates/*.gohtml"),
)
