package main

import "fmt"

func main() {

	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15
	fmt.Println("len - ", len(customerNumber))
	fmt.Println("cap - ", cap(customerNumber))

	customerNumber = append(customerNumber, 17)
	fmt.Println("len - ", len(customerNumber))
	fmt.Println("cap - ", cap(customerNumber))

	customerNumber = append(customerNumber, 19)
	customerNumber = append(customerNumber, 21)
	customerNumber = append(customerNumber, 22)
	customerNumber = append(customerNumber, 23)
	fmt.Println("len - ", len(customerNumber))
	fmt.Println("cap - ", cap(customerNumber))

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])
	fmt.Println(customerNumber[3])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
