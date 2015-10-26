package main

import "fmt"

func main() {

	// Maps - Shorthand Notation
	myGreeting := map[string]string{
		"Tim":     "Good morning!",
		"Jenny":   "Bonjour!",
		"Medhi":   "Buenos dias!",
		"Marcus":  "Bongiorno!",
		"Julian":  "Ohayo!",
		"Sushant": "Selamat pagi!",
		"Jose":    "Gutten morgen!",
	}

	fmt.Println(myGreeting["Jenny"])
}
