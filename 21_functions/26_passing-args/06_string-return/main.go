package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // 0x8201e0300
	fmt.Println(name) // Todd

	name = changeMe(name)

	fmt.Println(&name) //0x8201e0300
	fmt.Println(name) //Apollo
}

func changeMe(x string) string {
	fmt.Println(&x) // 0x8201e0320
	fmt.Println(x) // Todd
	x = "Apollo"
	fmt.Println(&x) // 0x8201e0320
	fmt.Println(x) // Apollo
	return x
}
