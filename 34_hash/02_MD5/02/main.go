package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// one way
	// reads bytes on the fly
	f, err := os.Open("../resources/table.csv")
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	fmt.Printf("The hash (sum) is: %x\n", h.Sum(nil))

	// or this way
	// but reads all the bytes at once, then does it
	f.Seek(0, 0)
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("readall didn't read well", err.Error())
	}
	fmt.Printf("The hash (sum) is: %x\n", md5.Sum(bs))
}
