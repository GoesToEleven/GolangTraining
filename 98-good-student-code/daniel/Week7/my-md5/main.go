package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func walkPath(path string, info os.FileInfo, err error, wg *sync.WaitGroup) {
	if !info.IsDir() {
		f, err := os.Open(path)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer f.Close()
		hash := md5.New()
		io.Copy(hash, f)
		fmt.Printf("%x\t%s\n", hash.Sum(nil), info.Name())
	}
	wg.Done()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please pass in a filename to calculate the md5 of")
	}

	var wg sync.WaitGroup
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		wg.Add(1)
		go walkPath(path, info, err, &wg)
		return nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	wg.Wait()
}
