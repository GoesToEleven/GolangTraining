package main

import "fmt"

func main() {

	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Todd"]) // 44
}

func changeMe(x map[string]int) {
	x["Todd"] = 44
}
