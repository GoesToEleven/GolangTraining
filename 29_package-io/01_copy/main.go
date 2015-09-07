package main

import (
	"strings"
	"io"
	"os"
)

func main() {
	rdr := strings.NewReader("test")
	io.Copy(os.Stdout, rdr)
}
