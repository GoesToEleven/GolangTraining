package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	myGreeting["key"] = "value"
	fmt.Println(myGreeting)
	fmt.Println(myGreeting["key"])
}
