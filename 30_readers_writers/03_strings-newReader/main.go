package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {

	rdr := strings.NewReader("some string")

	bs, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatalln("my program broke")
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}