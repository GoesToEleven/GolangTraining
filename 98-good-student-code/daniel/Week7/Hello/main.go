package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Print("What is your name? ")
	var myName string
	fmt.Scanf("%s", &myName)
	fmt.Println("It is nice to meet you", myName)
}
