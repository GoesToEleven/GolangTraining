package main

import "fmt"

type Vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type Car struct {
	Vehicle
	Wheels int
	Doors  int
}

type Plane struct {
	Vehicle
	Jet bool
}

type Boat struct {
	Vehicle
	Length int
}

func (v Vehicle) Specs() {
	fmt.Printf("Seats %v, max speed %v, color %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	prius := Car{Vehicle{6, 120, "white"}, 4, 5}
	prius.Specs()
}
