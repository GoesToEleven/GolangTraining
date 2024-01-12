package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// walk location
	root := "."

	directories, err := Directories(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create MD file
	file, err := os.OpenFile("INDEX.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	_, _ = file.WriteString("# Directory Index\n\n")

	// walk directories
	for _, directory := range directories {
		_ = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				depth := strings.Count(path, string(filepath.Separator)) - strings.Count(root, string(filepath.Separator))
				indent := strings.Repeat("- ", depth+1)
				_, _ = fmt.Fprintf(file, "%s[%s](%s)\n", indent, info.Name(), path)
			}
			return nil
		})
	}

	fmt.Println("directory index structure created")
}

// Directories List of level 1
func Directories(root string) ([]string, error) {
	directories := make([]string, 0)
	entries, err := os.ReadDir(root)
	if err != nil {
		return directories, err
	}
	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			directories = append(directories, e.Name())
		}
	}
	return directories, nil
}
