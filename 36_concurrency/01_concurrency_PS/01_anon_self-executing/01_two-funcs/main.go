package main

import "fmt"

func main() {

	func() {
		fmt.Println("One")
	}()

	func() {
		fmt.Println("Two")
	}()
}
