package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}
	myMap := map[string]string{
		"key1": "value",
		"key2": "value2",
	}

	fmt.Println(myGreeting["Jenny"])
	fmt.Println(myMap)
}
