package main

import "fmt"

func main() {
	intro := "Four score and seven years ago...."
	fmt.Println(intro)
	fmt.Println(&intro)
	intro = "Hahahaha!"
	fmt.Println(intro)
	fmt.Println(&intro)
	//  the below is invalid
	//	intro[0] = 70
	//	fmt.Println(intro)
	//	fmt.Println(&intro)
}
