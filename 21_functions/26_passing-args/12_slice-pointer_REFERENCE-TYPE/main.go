package main

import "fmt"

func main() {
	m := make([]string, 5, 25)
	fmt.Println(m) // [    ]
	changeMe(m)
	fmt.Println(m) // [Todd Rio   ]
}

func changeMe(x []string) {
	x[0] = "Todd"
	x[1] = "Rio"
	fmt.Println(x) // [Todd Rio   ]
}

