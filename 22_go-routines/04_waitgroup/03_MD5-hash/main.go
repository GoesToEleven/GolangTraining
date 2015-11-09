package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func md5file(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	fmt.Printf("%s %x\n", fileName, h.Sum(nil))
	wg.Done()
}

func walk() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go md5file(path)
		return nil
	})
}

func main() {
	walk()
	wg.Wait()
}
