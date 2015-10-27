package main

import "fmt"

func main() {
	var smaller, larger int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&smaller)
	fmt.Print("Enter a second larger number: ")
	fmt.Scanln(&larger)
	fmt.Println(larger % smaller)
}
