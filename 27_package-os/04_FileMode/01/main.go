package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.ModeDir)
	fmt.Printf("%p\n", os.ModeDir)
	fmt.Printf("%d\n", os.ModeDir)
	fmt.Println(os.ModeAppend)
	fmt.Println(os.ModeExclusive)
	fmt.Println(os.ModeTemporary)
	fmt.Println(os.ModeSymlink)
	fmt.Println(os.ModeDevice)
	fmt.Println(os.ModeNamedPipe)
	fmt.Println(os.ModeSocket)
	fmt.Println(os.ModeSetuid)
	fmt.Println(os.ModeSetgid)
	fmt.Println(os.ModeCharDevice)
	fmt.Println(os.ModeSticky)
	fmt.Println(os.ModeType)
	fmt.Println(os.ModePerm)
}
