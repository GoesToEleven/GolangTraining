package main

import (
	"fmt"
	"time"
)

func main() {
	timeAsString := "01/22/2012"
	fmt.Println(time.Parse("01/02/2006", timeAsString))
}
