package stripeexample

import (
	"net/http"

	"fmt"
	"html/template"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.gohtml"))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/payment", handlePayment)

}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/minions.jpg" {
		http.ServeFile(res, req, "minions.jpg")
		return
	}
	tpls.ExecuteTemplate(res, "index.gohtml", nil)
}

func handlePayment(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	stripeToken := req.FormValue("stripeToken")
	fmt.Fprintln(res, stripeToken)

}
