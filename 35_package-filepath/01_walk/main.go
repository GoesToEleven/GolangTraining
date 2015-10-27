package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walk() {
	var counter int
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		counter++
		fmt.Println(path)
		return nil
	})
	fmt.Println(counter)
}

func main() {
	walk()
}
