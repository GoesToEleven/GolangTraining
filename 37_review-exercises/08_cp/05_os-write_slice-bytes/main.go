package main

import (
	"log"
	"os"
)

func main() {

	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	dst.Write([]byte("Hello World"))
}

/*

os.Create
func Create(name string) (*File, error)

os
func (f *File) Write(b []byte) (n int, err error)

os
func (f *File) Read(b []byte) (n int, err error)

*/

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
programName initial.txt second.txt

*/
