package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"James", 20}
	fmt.Println(p1.name)
	p1.name = "Ian"
	fmt.Println(p1.name)
}
