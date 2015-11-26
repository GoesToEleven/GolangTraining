package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
