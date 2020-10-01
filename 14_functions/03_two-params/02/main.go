package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname, lname string) string {
	fmt.Println(fname, lname)
}
