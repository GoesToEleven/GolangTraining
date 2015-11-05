package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)

}