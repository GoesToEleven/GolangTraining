package main

import "fmt"

func main() {
	myConditional(false)
}

func myConditional(b bool) {

	if b {
		fmt.Println("first statement ran - ", b)
	}

	if !b {
		fmt.Println("second statement ran - ", b)
	}

}
