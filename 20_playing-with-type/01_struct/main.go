package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"James", 20}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Printf("%T\n", p1)
}

// we've already seen the above code
