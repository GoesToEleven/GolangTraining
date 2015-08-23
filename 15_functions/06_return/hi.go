package main

import "fmt"

func greet(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}