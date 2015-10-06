package main
import (
	"os"
	"path/filepath"
	"fmt"
)

func main() {
	wd, _ := os.Getwd()
	imgDir := "assets/imgs"
	filename := "01.jpg"
	path := filepath.Join(wd, imgDir, filename)
	fmt.Println(path)
}
