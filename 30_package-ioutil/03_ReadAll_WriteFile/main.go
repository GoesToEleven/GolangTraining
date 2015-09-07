package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	myStr := "Hello Everybody"

	err := ioutil.WriteFile("hey.txt", []byte(myStr), 0777)
	if err != nil {
		log.Fatalln("something went wrong!", err.Error())
	}

	f, err := os.Open("hey.txt")
	if err != nil {
		log.Fatalln("couldn't read file!", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Readall crashed!", err.Error())
	}

	fmt.Println(string(bs))

}
