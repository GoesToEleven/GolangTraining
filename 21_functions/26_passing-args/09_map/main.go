package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	fmt.Println(&m)
	m["Bob"] = 44
	fmt.Println(m["Bob"])
}

