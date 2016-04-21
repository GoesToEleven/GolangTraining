package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}

func main() {
	email := os.Args[1]
	gravatarHash := getGravatarHash(email)
	fmt.Println(`<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <img src="http://www.gravatar.com/avatar/` + gravatarHash + `?d=monsterid&f=1">
  </body>
</html>`)
}
