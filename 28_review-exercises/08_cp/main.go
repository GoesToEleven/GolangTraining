package main

import (
	"os"
	"log"
	"io/ioutil"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]

	f, err := os.Open(src)
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer f.Close()

	nf, err := os.Create(dst)
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke reading: ", err.Error())
	}

	_, err = nf.Write(bs)
	if err != nil {
		log.Fatalln("my program broke writing: ", err.Error())
	}
}

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
08_cp initial.txt second.txt

*/