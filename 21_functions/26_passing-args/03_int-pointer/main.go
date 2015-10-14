package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0x82023c080

	changeMe(&age)

	fmt.Println(&age) //0x82023c080
	fmt.Println(age) //24
}

func changeMe(x *int) {
	fmt.Println(x) // 0x82023c080
	fmt.Println(*x) // 44
	*x = 24
	fmt.Println(x) // 0x82023c080
	fmt.Println(*x) // 24
}
