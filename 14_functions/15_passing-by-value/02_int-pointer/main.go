package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0x82023c080

	changeMe(&age)

	fmt.Println(&age) //0x82023c080
	fmt.Println(age)  //24
}

func changeMe(z *int) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 24
}
