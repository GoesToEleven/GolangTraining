package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0x820252220

	changeMe(age)

	fmt.Println(&age) //0x820252220
}

func changeMe(x int) {
	fmt.Println(&x) // 0x820252240
	x = 24
	fmt.Println(&x) // 0x820252240
}
