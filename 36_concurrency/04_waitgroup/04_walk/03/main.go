package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func walkStep(dir string, fileNameChannel chan<- string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fileNameChannel <- path

		return nil
	})
	close(fileNameChannel)
}

func md5Step(fileNameChannel <-chan string, sumInfoChannel chan<- string) {
	for fileName := range fileNameChannel {
		bs := md5file(fileName)
		sumInfoChannel <- fmt.Sprintf("%s %x", fileName, bs)
	}
	close(sumInfoChannel)
}

func printStep(sumInfoChannel <-chan string) {
	for sumInfo := range sumInfoChannel {
		fmt.Println(sumInfo)
	}
}

func main() {
	fileNameChannel, sumInfoChannel := make(chan string), make(chan string)
	go walkStep(".", fileNameChannel)
	go md5Step(fileNameChannel, sumInfoChannel)
	printStep(sumInfoChannel)

}


