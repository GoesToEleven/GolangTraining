package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func main() {
	s := square{10}
	fmt.Println("Area: ", s.area())
}
