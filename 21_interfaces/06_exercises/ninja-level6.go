package main

import "fmt"

func main() {
	solution3()
	solution4()
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
