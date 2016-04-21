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
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		text := `


/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/`

		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
		return nil
	})
}
