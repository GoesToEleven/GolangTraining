package main

import (
	"io/ioutil"
)

func main() {

	err := ioutil.WriteFile("hello.txt", []byte("Hello world"), 0777)
	if err != nil {
		panic("something went wrong")
	}
}
