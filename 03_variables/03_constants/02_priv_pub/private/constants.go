package main

import (
	"fmt"
	"github.com/goestoeleven/GolangTraining/03_variables/03_constants/02_priv_pub/public"
)

const p string = "death & taxes"

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(public.ThisIsPublic)
	public.SomeFunc()
	// the below won't run
	// fmt.Println(public.thisIsNotPublic)
}