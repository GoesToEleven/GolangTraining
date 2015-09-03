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
	prius := Car{}
	tacoma := Car{}
	bmw528 := Car{}
	cars := []Car{prius, tacoma, bmw528}

	boeing747 := Plane{}
	boeing757 := Plane{}
	boeing767 := Plane{}
	planes := []Plane{boeing747, boeing757, boeing767}

	sanger := Boat{}
	nautique := Boat{}
	malibu := Boat{}
	boats := []Boat{sanger, nautique, malibu}

	for key, value := range cars {
		fmt.Println(key, " - ", value)
	}

	for key, value := range planes {
		fmt.Println(key, " - ", value)
	}

	for key, value := range boats {
		fmt.Println(key, " - ", value)
	}
}
