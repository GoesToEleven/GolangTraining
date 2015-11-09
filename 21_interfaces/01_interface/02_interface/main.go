package main

import "fmt"

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

func main() {
	s := Square{10}
	info(s)
}