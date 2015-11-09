package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// another shape
type Circle struct {
	radius float64
}
// which implements the Shape interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}