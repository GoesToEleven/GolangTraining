package main

import "fmt"

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}

func main() {
	greet("Jane", "Doe")
}