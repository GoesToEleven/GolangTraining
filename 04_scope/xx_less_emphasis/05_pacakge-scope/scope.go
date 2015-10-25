package main

import "fmt"

var x string = "Boss Hogg"

func oneFunc() {
	fmt.Println(x)
}

func twoFunc() {
	fmt.Println(x)
}

func main() {
	oneFunc()
	twoFunc()
}
