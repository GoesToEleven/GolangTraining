package main

import (
	"fmt"
	"math"
)

func main() {
	solution3()
	solution4()
	solution5()
	solution6()
	solution7()
	fmt.Println(solution8()())
}

func solution3() {

	y := []int{10, 200, 30}
	z := []int{10, 200, 300}


	defer fmt.Println(foo(y...))
	fmt.Println(bar(z))
}

func foo(x ...int) int {

	sum := 0
	for _, num := range x {
		sum += num
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, num := range x {
		sum += num
	}
	return sum
}


//Solution 4

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("\n name: %s %s age: %d", p.first, p.last, p.age)

}

func solution4() {

	p1 := person{
		"abk",
		"ku",
		41,
	}
	p1.speak()
}


type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s.area())
}

func (s Square) area() float64 {
	return (s.length * s.width)
}

func (c Circle) area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func solution5() {

	s1 := Square{10, 10}
	c1 := Circle{9}
	info(s1)
	info(c1)
}


func solution6() {

	func(x int) {
		fmt.Println("I am anonymous function", x)
	}(10)
}

func solution7() {
	f1 := func() {
		fmt.Println("Function Assignment ")
	}
	f1()
}

func solution8() func() int {

	return func() int {
			return 42
	}
}