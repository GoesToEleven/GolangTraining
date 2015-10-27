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

func measure(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	circ := Circle{5}
	sqr := Square{10}
	measure(circ)
	measure(sqr)
}
