package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11,}

	for i, value := range mySlice {
		fmt.Println(i, " - ", value)
	}

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 13)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

}