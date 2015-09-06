package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x", x)
	fmt.Println("y", y)
}

piazza.com/fresnostate/fall2015/csci191t