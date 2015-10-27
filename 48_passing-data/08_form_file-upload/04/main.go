package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.ListenAndServe(":9000", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			file, _, err := req.FormFile("my-file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer file.Close()

			src := io.LimitReader(file, 400)

			dst, err := os.Create(filepath.Join(".", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()

			io.Copy(dst, src)
		}

		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `
      <form method="POST" enctype="multipart/form-data">
        <input type="file" name="my-file">
        <input type="submit">
      </form>
      `)
	}))
}
