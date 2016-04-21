package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name(), info.Size(), info.Mode(), info.IsDir())
		return nil
	})
}

/*
walk is recursive
readdir is not
*/
