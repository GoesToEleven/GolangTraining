package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func cp(srcName, dstName string) error {

	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("my program broke opening: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("my program broke creating:%v ", err)
	}

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		return fmt.Errorf("my program broke reading: %v", err)
	}

	_, err = dst.Write(bs)
	if err != nil {
		return fmt.Errorf("my program broke writing: %v", err)
	}

	return nil
}

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("Usage: 03 <SRC> <DST>")
	}

	srcName := os.Args[1]
	dstName := os.Args[2]

	err := cp(srcName, dstName)
	if err != nil {
		log.Fatalln(err)
	}

}

/*

step 1 - at command line enter:
go install

step 2 - at command line enter:
03 initial.txt second.txt

*/