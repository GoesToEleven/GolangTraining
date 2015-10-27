package hello

import "fmt"

func ByeBye() {
	fmt.Println(looper("Bye! "))
}

func looper(str string) string {
	var newString string
	x := 7
	for i := 0; i < 100; i++ {
		newString += str
		x := 9
		fmt.Println(x)
	}
	fmt.Println(x)
	return newString
}
