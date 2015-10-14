package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // 0x820250220

	changeMe(name)

	fmt.Println(&name) //0x820250220
}

func changeMe(x string) {
	fmt.Println(&x) // 0x820250240
	x = "Apollo"
	fmt.Println(&x) // 0x820250240
}
