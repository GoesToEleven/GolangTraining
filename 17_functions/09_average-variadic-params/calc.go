package main

import "fmt"

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}
