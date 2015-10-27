package main

import "fmt"

func main() {
	fmt.Println(fibo(10))
}

func fibo(xs int) int {
	if xs == 1 {
		return 1
		fmt.Println("xs == 1, return 1")
	} else if xs == 0 {
		return 0
		fmt.Println("xs == 0, return 0")
	}
	fmt.Println(xs, "...so...", xs-1, xs-2)
	return fibo(xs-1) + fibo(xs-2)
}
