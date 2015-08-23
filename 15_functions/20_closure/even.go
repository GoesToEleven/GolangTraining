package main

import "fmt"

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	fmt.Println(nextEven()) // 6
}
