package main

import "fmt"

func main() {

	go func() {
		fmt.Println("One")
	}()

	go func() {
		fmt.Println("Two")
	}()
}
