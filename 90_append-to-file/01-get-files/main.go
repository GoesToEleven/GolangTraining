package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, ".git") ||
		strings.Contains(path, ".idea") ||
		strings.Contains(path, ".DS_Store") {
			return nil
		}

		fmt.Println(path)
		return nil
	})
}
