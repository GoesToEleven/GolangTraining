package main

import (
	"os"
	"log"
	"io/ioutil"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("Usage: 02 <SRC> <DST>")
	}

	srcName := os.Args[1]
	dstName := os.Args[2]

	src, err := os.Open(srcName)
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("my program broke reading: ", err.Error())
	}

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("my program broke writing: ", err.Error())
	}
}

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
02 initial.txt second.txt

*/