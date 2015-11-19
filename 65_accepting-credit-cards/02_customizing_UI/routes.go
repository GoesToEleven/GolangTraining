package stripeexample

import (
	"net/http"

	"html/template"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/minions.jpg" {
		http.ServeFile(res, req, "minions.jpg")
		return
	}
	tpls.ExecuteTemplate(res, "index.gohtml", nil)
}
