package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	sum(aSlice...)
	sum()
}

func sum(numbers ...int) {
	fmt.Println("here are the arguments: ", numbers)
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println("here is the total: ", total)
	fmt.Println()
}