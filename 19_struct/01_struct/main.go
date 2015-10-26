package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"James", 20})
}
