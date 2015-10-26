package main

import "fmt"

func main() {
	intro := "Four 世"
	fmt.Printf("%T\n", intro)
	fmt.Println(intro)
	bs := []byte(intro)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println("*********")
	fmt.Printf("%d\n", bs)

	for _, v := range bs {
		fmt.Printf("%d\t\t %#x\t %b\n", v, v, v)
	}
	fmt.Println("*********")
	y := 9999999999999999

	fmt.Printf("%d\t\t %#x\t %b\n", y, y, y)
	fmt.Println(&y)
	fmt.Sprint(y)
	fmt.Println("*********")

	z := 'h'
	fmt.Printf("%T\n", z)
}
