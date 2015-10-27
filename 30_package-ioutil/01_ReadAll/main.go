package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke again")
	}

	str := string(bs)
	fmt.Println(str)
}

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
06_cat main.go

*/
