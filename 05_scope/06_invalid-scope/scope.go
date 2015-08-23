package main

import "fmt"

func oneFunc() {
	var x string = "Boss Hogg"
	fmt.Println(x)
}

func twoFunc() {
	fmt.Println(x)
}

func main() {
	oneFunc()
	twoFunc()
}
