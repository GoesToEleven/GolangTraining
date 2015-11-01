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
	info(c)
	info(s)

}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

