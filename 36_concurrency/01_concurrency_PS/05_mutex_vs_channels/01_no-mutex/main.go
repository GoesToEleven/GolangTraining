package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
			}()
		}
	}

	fmt.Scanln()
}
