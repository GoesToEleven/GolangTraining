package main
import (
	"os"
	"path/filepath"
	"fmt"
)

func main() {
	wd, _ := os.Getwd()
	filename := "01.jpg"
	path := filepath.Join(wd, "assets", "imgs", filename)
	fmt.Println(path)
}
