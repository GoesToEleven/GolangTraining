package main

import "fmt"

func main() {
	f(1, 2)
	f(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	f(aSlice...)
	f()
}

func f(numbers ...int) {
	fmt.Println(numbers)
}