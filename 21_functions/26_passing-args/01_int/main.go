package main

import "fmt"

func main() {

	age := 44

	changeMe(age)

	fmt.Println(age) //44
}

func changeMe(x int) {
	fmt.Println(x) // 44
	x = 24
	fmt.Println(x) // 24
}
