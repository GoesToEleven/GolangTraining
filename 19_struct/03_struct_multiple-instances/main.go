package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"James", 20}
	p2 := person{"Ian", 45}
	fmt.Println(p1.name)
	fmt.Println(p2.name)
}
