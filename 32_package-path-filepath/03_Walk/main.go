package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		fmt.Println(info)
		return nil
	})
}

/*
walk is recursive
readdir is not
*/
