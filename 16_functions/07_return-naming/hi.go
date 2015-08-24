package main

import "fmt"

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

// we can give a name to the return type