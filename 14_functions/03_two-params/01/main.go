package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
