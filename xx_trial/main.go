package main

import "fmt"

func main() {

	m := "Hello World"
	n := &m
	fmt.Println(m, n)

//	fmt.Println("m - ", m)
//	fmt.Println([]byte(m))
//	fmt.Println([]byte(m)[0])
//	fmt.Printf("%T\n",[]byte(m)[0])
//	fmt.Printf("%b\n",[]byte(m)[0])
//	fmt.Println()
//	for _, v := range m {
//		fmt.Printf("%b\n",v)
//	}
}