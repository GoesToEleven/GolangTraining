package main

import "fmt"

func main() {
	var age *int = new(int) // new returns a pointer
	fmt.Println(age)
	fmt.Println(*age)
}
