package main

import "fmt"

type person struct {
	naem string
	aeg int
}

func main() {
	p1 := new(person)
	p1.naem = "kai"
	fmt.Println(p1)
}
