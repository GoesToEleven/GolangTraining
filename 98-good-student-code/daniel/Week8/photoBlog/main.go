package photo

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type fileTimes []time.Time

func (ft *fileTimes) Len() int {
	return len(*ft)
}

func (ft *fileTimes) Less(i, j int) bool {
	return (*ft)[i].After((*ft)[j])
}

func (ft *fileTimes) Swap(i, j int) {
	(*ft)[i], (*ft)[j] = (*ft)[j], (*ft)[i]
}

func mainSite(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	tpl, err := template.ParseFiles("mainSite.gohtml")
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
	times := fileTimes{}
	images := map[time.Time]string{}
	filepath.Walk("images", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileTime := info.ModTime()
			images[fileTime] = filepath.ToSlash("/" + path)
			times = append(times, fileTime)
		}
		return nil
	})
	sort.Sort(&times)
	sortedImages := make([]string, len(times))
	for i, v := range times {
		sortedImages[i] = images[v]
	}
	err = tpl.Execute(res, sortedImages)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}

func adminSite(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	gotFile := false
	if req.Method == "POST" {
		file, _, err := req.FormFile("image")
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}
		defer file.Close()

		ckSum := md5.New()
		io.Copy(ckSum, file)
		filename := "images/" + hex.EncodeToString(ckSum.Sum(nil))
		wtr, err := os.Create(filename)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}
		defer wtr.Close()

		_, err = file.Seek(0, 0)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}

		_, err = io.Copy(wtr, file)
		if err != nil {
			http.Error(res, "Server Error", http.StatusInternalServerError)
			log.Errorf(ctx, err.Error())
			return
		}

		gotFile = true
	}

	tpl, err := template.ParseFiles("adminSite.gohtml")
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}

	err = tpl.Execute(res, gotFile)
	if err != nil {
		http.Error(res, "Server Error", http.StatusInternalServerError)
		log.Errorf(ctx, err.Error())
		return
	}
}

func getCSS(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "style.css")
}

func init() {
	imagesHandler := http.StripPrefix("/images/", http.FileServer(http.Dir("images/")))

	http.HandleFunc("/", mainSite)
	http.HandleFunc("/admin", adminSite)
	http.Handle("/images/", imagesHandler)
	http.HandleFunc("/style.css", getCSS)
}
