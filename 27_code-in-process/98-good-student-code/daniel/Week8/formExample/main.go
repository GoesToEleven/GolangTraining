package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleForm(res http.ResponseWriter, req *http.Request) {
	first := req.FormValue("first")
	last := req.FormValue("last")

	rdr, hdr, err := req.FormFile("file")
	if rdr != nil {
		if err != nil {
			http.Error(res, "Error reading file", 500)
			fmt.Println(err)
			return
		}
		defer rdr.Close()
		wtr, err := os.Create(hdr.Filename)
		if err != nil {
			http.Error(res, "Error writing file", 500)
			return
		}
		defer wtr.Close()
		io.Copy(wtr, rdr)
	}

	res.Header().Set("Content-Type", "text/html")
	if first != "" && last != "" {
		fmt.Fprintf(res, "<p>Hello %s %s</p>", first, last)
	}
	io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
  <input type="text" name="first" placeholder="First name" required>
  <input type="text" name="last" placeholder="Last name" required>
  <input type="file" name="file">
  <input type="submit">
</form>`)
}

func main() {
	http.HandleFunc("/", handleForm)
	http.ListenAndServe(":9000", nil)
}
