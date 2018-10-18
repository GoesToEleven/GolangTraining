package main

import (
"fmt"
)

type person struct {
	fname string
	lname string
	ficf  []string
}

func main() {

	fmt.Println("\nSolution 2\n")
	solution2()

	fmt.Println("\nSolution 3\n")
	solution3()

	fmt.Println("\nSolution 4\n")
	solution4()
}

func solution2() {

	//Creation of maps using structs
	p1 := person{
		fname: "aj",
		lname: "kul",
		ficf:  []string{"kasata", "vanilla"},
	}
	p2 := person{
		fname: "Ro",
		lname: "kul1",
		ficf:  []string{"rainbow", "mint", "chocolate"},
	}

	/*
		fmt.Println(p1)
		for _, icecream := range p1.ficf {
			fmt.Println(icecream)
		}
		for _, icecream := range p2.ficf {
			fmt.Println(icecream)
		}
	*/

	m := map[string]person{
		p1.lname: p1,
		p2.lname: p2,
	}

	fmt.Println(m["kul"])
}


type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func solution3() {

	// Usage of embedded structs
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red"},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Printf("%d %s %t", t1.doors, t1.color, t1.fourWheel)


}



func solution4() {

	// Usage of anonymous struct
	t1 := struct {
		doors int
		color string
	}{
		doors: 4,
		color: "red",
	}
	fmt.Println(t1)
}