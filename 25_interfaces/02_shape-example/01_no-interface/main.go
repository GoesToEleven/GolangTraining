package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.side * s.side
}

func main() {
	c := Circle{5}
	s := Square{10}
	totalArea := c.area() + s.area()
	fmt.Println("Total Area: ", totalArea)
}

// what if I had thousands of shapes?
// how would I create a function to sum their areas?
