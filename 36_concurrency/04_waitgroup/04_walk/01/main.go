package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func md5file(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	return h.Sum(nil)
}

func main() {
	var wg sync.WaitGroup

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go func() {
			bs := md5file(path)
			fmt.Printf("%s %x\n", path, bs)
			wg.Done()
		}()
		return nil
	})
	wg.Wait()
}


