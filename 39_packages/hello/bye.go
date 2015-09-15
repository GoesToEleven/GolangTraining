package hello

import "fmt"

func ByeBye() {
	fmt.Println(looper("Bye! "))
}

func looper(str string) string {
	var newString string
	for i := 0; i < 100; i++ {
		newString += str
	}
	return newString
}