package main

import "fmt"

func main() {
	var x int = 5
	str := "Hello world " + fmt.Sprint(x)
	fmt.Println(str)
}
