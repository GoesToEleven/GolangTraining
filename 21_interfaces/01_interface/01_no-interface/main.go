package main

import (
	"fmt"
)

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func main() {
	s := Square{10}
	fmt.Println("Area: ", s.area())
}