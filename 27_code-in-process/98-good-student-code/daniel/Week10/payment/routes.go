package payment

import (
	"html/template"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	r := httprouter.New()
	r.GET("/", handleIndex)
	r.POST("/payment", handlePayment)
	http.Handle("/", r)
}

func handlePayment(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	token := req.FormValue("stripeToken")
	ctx := appengine.NewContext(req)
	err := chargeAccount(ctx, token)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(res, "Thanks for the money")
}

func handleIndex(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "index", nil)
	if err != nil {
		ctx := appengine.NewContext(req)
		http.Error(res, "Server error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}
