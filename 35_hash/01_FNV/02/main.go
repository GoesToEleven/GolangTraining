package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := fnv.New64()
	io.Copy(h, f)
	fmt.Println("The sum is:", h.Sum64())
}
