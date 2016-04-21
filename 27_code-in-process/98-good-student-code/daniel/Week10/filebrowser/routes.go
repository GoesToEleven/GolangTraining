package browser

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
	r := httprouter.New()
	r.GET("/", handleCredentials)
	r.POST("/", handleLogin)
	r.GET("/browse/*path", handlePath)
	http.Handle("/", r)
}
