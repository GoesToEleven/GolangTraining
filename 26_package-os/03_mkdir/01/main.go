package main

import (
	"log"
	"os"
)

func main() {

	err := os.Mkdir("/somefolderthatdoesntexist", 0x777)
	if err != nil {
		log.Fatalln("my program broke on mkdir: ", err.Error())
	}

	f, err := os.Create("/somefolderthatdoesntexist/hello.txt")
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	str := "Put some phrase here."
	bs := []byte(str)

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
sudo 05_mkdir_write-file_absolute-path

--- or ---

step 1 - at command line enter:
sudo go run main.go

---------

use at command line to see folder:
ls /somefolderthatdoesntexist

use at command line to remove folder:
rm -rf /somefolderthatdoesntexist

*/
