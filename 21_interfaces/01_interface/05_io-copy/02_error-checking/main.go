package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	if err != nil {
		fmt.Println(err)
		return
	}

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Get("http://www.mcleods.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		fmt.Println(err)
	}

}
