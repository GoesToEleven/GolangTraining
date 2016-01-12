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

		if !strings.Contains(path, ".go") {
			return nil
		}

		fmt.Println(path)
		return nil

	})
}
