package main

import "fmt"

func main() {
	Greeting("hello: ", "James", "Jasmine")
}

func Greeting(prefix string, who ...string) {
	fmt.Println(prefix)
	for _, value := range who {
		fmt.Println(value)
	}
}