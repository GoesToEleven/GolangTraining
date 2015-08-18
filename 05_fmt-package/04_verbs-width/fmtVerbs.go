package main

import "fmt"

func main() {

	a := 762324.84274232380972

	fmt.Printf("%%v    - %v\n", a)
	fmt.Printf("%%f    - %f\n", a)
	fmt.Printf("%%.2f  - %.2f\n", a)
	fmt.Printf("%%9.2f - %9.2f\n", a)
	fmt.Printf("%%4.3f - %4.3f\n", a)
	fmt.Printf("%%20.5f - %20.5f\n", a)
	fmt.Printf("%%9.f  - %9.f\n", a)
}