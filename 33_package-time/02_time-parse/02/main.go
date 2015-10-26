package main

import (
	"fmt"
	"time"
)

func main() {
	timeAsString := "01/22/2012"
	fmt.Println(time.Parse("01/01_this-does-not-compile/2006", timeAsString))
}
