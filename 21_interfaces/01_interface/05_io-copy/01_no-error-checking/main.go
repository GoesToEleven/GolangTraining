package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("http://www.mcleods.com")
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
