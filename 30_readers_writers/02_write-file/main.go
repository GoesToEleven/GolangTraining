package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	str := "Here is a phrase."
	bs := []byte(str)
	fmt.Println(str)
	fmt.Println(bs)

	f, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln("my program broke")
	}
	defer f.Close()

	n, err := f.Write(bs)
	if err != nil {
		log.Fatalln("my program broke")
	}
	fmt.Println(n)
}