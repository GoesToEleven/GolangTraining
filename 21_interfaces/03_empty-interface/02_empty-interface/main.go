package main

import "fmt"

type Vehicles interface{}

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
	prius := Car{}
	tacoma := Car{}
	bmw528 := Car{}
	boeing747 := Plane{}
	boeing757 := Plane{}
	boeing767 := Plane{}
	sanger := Boat{}
	nautique := Boat{}
	malibu := Boat{}
	rides := []Vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
