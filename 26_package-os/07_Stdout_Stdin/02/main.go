package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	rdr := strings.NewReader("test")
	io.Copy(os.Stdout, rdr)
}
