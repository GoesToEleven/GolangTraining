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

type Shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.side * s.side
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

// a new method which takes the INTERFACE TYPE Shape
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
	fmt.Println("Total Area: ", totalArea(c, s))
}
