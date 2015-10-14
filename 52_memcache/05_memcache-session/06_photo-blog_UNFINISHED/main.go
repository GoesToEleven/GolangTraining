package blog

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("assets/templates/*.html")

	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs", http.FileServer(http.Dir("./assets/imgs"))))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// get cookie UUID, or set
	cookie, _ := req.Cookie("sessionid")
	if cookie == nil {
		cookie = createCookie()
		http.SetCookie(res, cookie)
	}
	// get memcache session data
	item, _ := memcache.Get(ctx, cookie.Value)
	if item == nil {
		http.Redirect(res, req, "/login", 302)
		return
	}
	// unmarshal
	var m map[string]string
	json.Unmarshal(item.Value, &m)
	// authenticate
	if m["loggedin"] == "false" || m["loggedin"] == nil {
		http.Redirect(res, req, "/login", 302)
		return
	}
	// upload photo
	if req.Method == "POST" {
		src, hdr, err := req.FormFile("data")
		if err == nil {
			m = uploadPhoto(m, src, hdr)
		}
	}


	// save session
	session.Save(req, res)
	// get photos
	data := getPhotos(session)
	// execute template
	tpl.ExecuteTemplate(res, "index.html", data)
}

func createCookie() *http.Cookie {
	id, _ := uuid.NewV4()
	return &http.Cookie{
		Name:  "sessionid",
		Value: id.String(),
	}
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	session.Values["loggedin"] = "false"
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}

func login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" && req.FormValue("password") == "secret" {
		setMemcache(res, req)
		http.Redirect(res, req, "/", 302)
		return
	}
	// execute template
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func setMemcache(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// get cookie UUID, or set
	cookie, _ := req.Cookie("sessionid")
	if cookie == nil {
		cookie = createCookie()
		http.SetCookie(res, cookie)
	}
	m := map[string]string{
		"username": req.FormValue("userName"),
		"loggedin": "true",
		"photos": []string{},
	}
	bs, _ := json.Marshal(m)
	item := &memcache.Item{
		Key:   cookie.Value,
		Value: bs,
	}
	memcache.Set(ctx, item)
}

func uploadPhoto(m map[string]string, src multipart.File, hdr *multipart.FileHeader) map[string]string {
	defer src.Close()
	fName := getSha(src) + ".jpg"
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "assets", "imgs", fName)
	dst, _ := os.Create(path)
	defer dst.Close()
	src.Seek(0, 0)
	io.Copy(dst, src)
	addPhoto(m, fName)
	return m
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func addPhoto(m map[string]string, fName string) {
	data := getPhotos(m)
	data = append(data, fName)
	bs, _ := json.Marshal(data)
	m["data"] = string(bs)
}

func getPhotos(m map[string]string) []string {
	var data []string
	jsonData := m["data"]
	if jsonData != nil {
		json.Unmarshal([]byte(jsonData.(string)), &data)
	}
	return data
}
