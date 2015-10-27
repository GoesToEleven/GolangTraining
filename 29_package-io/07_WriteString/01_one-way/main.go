package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("hello.txt")
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

func WriteString(w Writer, s string) (n int, err error)

WriteString writes the contents of the string s to w,
which accepts a slice of bytes. If w implements a WriteString method, it is invoked directly.

*/
