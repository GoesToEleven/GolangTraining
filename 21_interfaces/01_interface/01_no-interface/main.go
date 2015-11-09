package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func main() {
	s := Square{10}
	fmt.Println("Area: ", s.area())
}