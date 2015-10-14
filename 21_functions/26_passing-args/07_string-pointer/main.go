package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name) //Apollo
}

func changeMe(x *string) {
	fmt.Println(x) // 0x82023c080
	fmt.Println(*x) // Todd
	*x = "Apollo"
	fmt.Println(x) // 0x82023c080
	fmt.Println(*x) // Apollo
}
