package movieinfo

import "net/http"

type indexModel struct {
}

var fileServer = http.FileServer(http.Dir("public/"))

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		fileServer.ServeHTTP(res, req)
		return
	}
	model := &indexModel{}
	err := tpl.ExecuteTemplate(res, "index", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}
