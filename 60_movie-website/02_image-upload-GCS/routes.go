package movieinfo

import (
	"net/http"
	"strings"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/new-movie", handleNewMovie)
	http.HandleFunc("/posters/", handlePosters)
}

func handlePosters(res http.ResponseWriter, req *http.Request) {
	posterID := strings.SplitN(req.URL.Path, "/", 3)[2]
	ctx := appengine.NewContext(req)
	// rc, err := getFile(ctx, posterID)
	// if err != nil {
	// 	http.NotFound(res, req)
	// 	return
	// }
	// defer rc.Close()
	//
	// io.Copy(res, rc)
	mediaLink, err := getFileLink(ctx, posterID)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	http.Redirect(res, req, mediaLink, 302)
}
