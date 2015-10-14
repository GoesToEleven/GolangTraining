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

	addEntry(myGreeting)

	fmt.Println(myGreeting["Harleen"])
}

func addEntry(m map[string]string) {
	m["Harleen"] = "Howdy"
	fmt.Println(m)
}
