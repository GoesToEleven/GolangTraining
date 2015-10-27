package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	myPhrase := "mmm, licorice"
	rdr := strings.NewReader(myPhrase)

	bs, err := ioutil.ReadAll(rdr)
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
