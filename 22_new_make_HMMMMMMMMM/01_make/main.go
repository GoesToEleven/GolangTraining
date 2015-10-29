package main

import "fmt"

func main() {

	favorite := make(map[string]string)
	fmt.Println(favorite)
	fmt.Println(len(favorite))
	favorite["breakfast"] = "eggs"
	fmt.Println(favorite)
	fmt.Println(len(favorite))
}
