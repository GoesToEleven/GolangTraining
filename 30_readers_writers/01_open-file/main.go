package main

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln("my program broke")
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke")
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}

// readall takes a reader
// file has a read method
// therefore file implements the reader interface
// and readall can take a file