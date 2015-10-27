package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func cp(srcName, dstName string) error {

	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	br := bufio.NewReader(src)

	_, err = io.Copy(dst, br)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v ", err)
	}

	return nil
}

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("Usage: 07_bufio_io-copy <SRC> <DST>")
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
programName initial.txt second.txt

*/
