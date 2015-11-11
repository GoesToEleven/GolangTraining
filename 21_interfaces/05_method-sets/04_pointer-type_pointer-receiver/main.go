package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s Shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := Circle{5}
	info(&c)
}

