package main

import "fmt"

func main() {

	// Maps - Shorthand Notation
	myGreeting := map[string]string{
		"Tim":     "Good morning!",
		"Jenny":   "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(len(myGreeting))
}
