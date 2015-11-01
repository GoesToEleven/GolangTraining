package browser

import (
	"io"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/julienschmidt/httprouter"
)

type browseModel struct {
	Path       string
	Bucket     string
	Subfolders []string
	Files      []string
}

func handlePath(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := appengine.NewContext(req)
	s := getSession(ctx, req)
	if s.Bucket == "" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	cctx, err := getCloudContext(ctx, s.Credentials)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	path := p.ByName("path")

	if !strings.HasSuffix(path, delimiter) {
		res.Header().Set("Content-Disposition", "attachment")
		f, err := getFile(cctx, s.Bucket, path[1:])
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}
		defer f.Close()
		io.Copy(res, f)
	} else {
		files, subfolders, err := listFiles(cctx, s.Bucket, path[1:])
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}
		data := browseModel{
			Path:       path,
			Bucket:     s.Bucket,
			Subfolders: subfolders,
			Files:      files,
		}
		err = tpl.ExecuteTemplate(res, "browse", data)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}
	}
}
