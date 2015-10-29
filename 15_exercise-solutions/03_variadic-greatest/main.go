package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(greatest)
}
