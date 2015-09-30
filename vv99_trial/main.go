package main

import "fmt"

func main () {
	str := greet()
	fmt.Println(str)
}

func greet() string {
	fmt.Print("What is your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("How old are you? ")
	var age int
	fmt.Scanln(&age)

	return fmt.Sprint(name, " is ", age, " years old")
}