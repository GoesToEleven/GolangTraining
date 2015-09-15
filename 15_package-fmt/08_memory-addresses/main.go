package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		fmt.Print(&i, " - ")
		fmt.Printf("%d\t", &i)
		fmt.Printf("%T\n", &i)
	}
}
