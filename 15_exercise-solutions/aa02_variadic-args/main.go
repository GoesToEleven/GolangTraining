package main

import "fmt"

func main() {
	s := []string{"James", "Jasmine"}
	Greeting("hello: ", s...)
}

func Greeting(prefix string, who ...string) {
	fmt.Println(prefix)
	for _, value := range who {
		fmt.Println(value)
	}
}