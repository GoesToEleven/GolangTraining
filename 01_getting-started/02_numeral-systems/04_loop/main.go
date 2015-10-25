package main

import "fmt"

func main() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
