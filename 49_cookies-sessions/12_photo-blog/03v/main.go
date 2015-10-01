package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"fmt"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))


func init() {
	var err error
	tpl, err = template.ParseFiles("assets/tpl/index.gohtml", "assets/tpl/admin_login.gohtml", "assets/tpl/admin_upload.gohtml")
	if err != nil {
		log.Fatalln("couldn't parse", err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/upload", upload)
	http.HandleFunc("/admin/logout", logout)
	http.Handle("/assets/imgs/", http.StripPrefix("/assets/imgs/", http.FileServer(http.Dir("assets/imgs/"))))
	http.ListenAndServe(":9000", context.ClearHandler(http.DefaultServeMux))
	/*
	go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
	 */
}

func home(res http.ResponseWriter, req *http.Request) {

	type Photo struct {
		PhotoPath string
		Lat float64
		Long float64
	}

	var model struct {
		Photos []Photo
		LoggedIn bool
	}

	// session
	session, _ := store.Get(req, "session-name")
	_, model.LoggedIn = session.Values["loggedin"]

	filepath.Walk("assets/imgs", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {

			fmt.Println("WALKING", path)

			var currentPhoto Photo

			currentPhoto.PhotoPath = path

			// photo lat long
//			fname := "assets/imgs/us.jpg"

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			// Optionally register camera makenote data parsing - currently Nikon and
			// Canon are supported.
			exif.RegisterParsers(mknote.All...)

			x, err := exif.Decode(f)
			if err != nil {
				log.Println("no info")
				return nil
			}

//			camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
//			fmt.Println(camModel.StringVal())
//
//			focal, _ := x.Get(exif.FocalLength)
//			numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
//			fmt.Printf("%v/%v", numer, denom)
//
//			// Two convenience functions exist for date/time taken and GPS coords:
//			tm, _ := x.DateTime()
//			fmt.Println("Taken: ", tm)

			currentPhoto.Lat , currentPhoto.Long, _ = x.LatLong()
			fmt.Println("lat, long: ", currentPhoto.Lat , ", ", currentPhoto.Long)

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
		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		session, _ := store.Get(req, "session-name")
		// Set some session values.
		session.Values["loggedin"] = "true"
		// Save it.
		session.Save(req, res)
		http.Redirect(res, req, "/admin/upload", 302)
		return
	}

	tpl.ExecuteTemplate(res, "admin_login.gohtml", nil)
}

func upload(res http.ResponseWriter, req *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(req, "session-name")
	// Set some session values.
	_, ok := session.Values["loggedin"]
	if !ok {
		//Change url
		http.Redirect(res, req, "/admin", 302)
		return
	}

	// if they are uploading a file, handle that
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

	// execute template
	tpl.ExecuteTemplate(res, "admin_upload.gohtml", ok)
}

func logout(res http.ResponseWriter, req *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(req, "session-name")
	delete(session.Values, "loggedin")
	session.Save(req, res)
	http.Redirect(res, req, "/admin", 302)
	return
}

/*
'https://maps.googleapis.com/maps/api/staticmap?center=' + e.coords.latitude + ','
                + e.coords.longitude + '&zoom=10&size=700x700&maptype=roadmap';


 */
