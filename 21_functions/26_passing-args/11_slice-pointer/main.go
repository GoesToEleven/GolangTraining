package main

import "fmt"

func main() {
	m := make([]string, 0, 25)
	fmt.Println(m)
	changeMe(&m)
	fmt.Println(m) //
}

func changeMe(x *[]string) {
	*x = append(*x, "Todd")
	*x = append(*x, "Rio")
	fmt.Println(*x)
}
