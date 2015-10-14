package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	fmt.Println(&m)
	m["Bob"] = 44
	fmt.Println(m["Bob"])
}

