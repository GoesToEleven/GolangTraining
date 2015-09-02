package main

import "fmt"

func main() {
	var name int
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
}