package photos

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	var err error
	tpl, err = template.ParseFiles("assets/tpl/index.gohtml", "assets/tpl/admin_login.gohtml", "assets/tpl/admin_upload.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse", err)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/upload", upload)
	http.HandleFunc("/admin/logout", logout)
	http.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs/", http.FileServer(http.Dir("assets/imgs/"))))
}

func home(res http.ResponseWriter, req *http.Request) {

	type Photo struct {
		PhotoPath string
		Lat       float64
		Long      float64
	}

	var model struct {
		Photos   []Photo
		LoggedIn bool
	}

	session, _ := store.Get(req, "session-name")
	_, model.LoggedIn = session.Values["loggedin"]

	filepath.Walk("assets/imgs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {

			fmt.Println("WALKING", path)

			var currentPhoto Photo

			currentPhoto.PhotoPath = path

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			exif.RegisterParsers(mknote.All...)

			x, err := exif.Decode(f)
			if err != nil {
				log.Println("no info")
				return nil
			}

			currentPhoto.Lat, currentPhoto.Long, _ = x.LatLong()
			fmt.Println("lat, long: ", currentPhoto.Lat, ", ", currentPhoto.Long)

			model.Photos = append(model.Photos, currentPhoto)
		}
		return nil
	})

	err := tpl.ExecuteTemplate(res, "index.gohtml", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func admin(res http.ResponseWriter, req *http.Request) {
	userName := req.FormValue("userName")
	password := req.FormValue("password")

	if userName == "You" && password == "Me" {
		session, _ := store.Get(req, "session-name")
		session.Values["loggedin"] = "true"
		session.Save(req, res)
		http.Redirect(res, req, "/admin/upload", 302)
		return
	}

	tpl.ExecuteTemplate(res, "admin_login.gohtml", nil)
}

func upload(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session-name")
	_, ok := session.Values["loggedin"]
	if !ok {
		http.Redirect(res, req, "/admin", 302)
		return
	}

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer src.Close()

		fileName := hdr.Filename
		dst, err := os.Create("assets/imgs/" + fileName)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
	}
	tpl.ExecuteTemplate(res, "admin_upload.gohtml", ok)
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session-name")
	delete(session.Values, "loggedin")
	session.Save(req, res)
	http.Redirect(res, req, "/admin", 302)
	return
}
