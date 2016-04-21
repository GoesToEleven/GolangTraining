package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, _ := os.Getwd()
	filename := "01.jpg"
	path := filepath.Join(wd, "assets", "imgs", filename)
	fmt.Println(path)
}
